#!/usr/bin/env bash

REGION=$1
STACK_NAME=$2

cd ./3-microservices

DEPLOYABLE_SERVICES=(
	identity
);

PRIMARY='\033[0;34m'
NC='\033[0m' # No Color

# Fetch the stack metadata for use later
printf "${PRIMARY}* Fetching current stack state${NC}\n";

QUERY=$(cat <<-EOF
[
	Stacks[0].Outputs[?OutputKey==\`ClusterName\`].OutputValue,
	Stacks[0].Outputs[?OutputKey==\`ALBArn\`].OutputValue,
	Stacks[0].Outputs[?OutputKey==\`ECSRole\`].OutputValue,
	Stacks[0].Outputs[?OutputKey==\`Url\`].OutputValue,
	Stacks[0].Outputs[?OutputKey==\`VPCId\`].OutputValue
]
EOF
)

RESULTS=$(aws cloudformation describe-stacks \
	--stack-name $STACK_NAME \
	--region $REGION \
	--query "$QUERY" \
	--output text);
RESULTS_ARRAY=($RESULTS)

CLUSTER_NAME=${RESULTS_ARRAY[0]}
ALB_ARN=${RESULTS_ARRAY[1]}
ECS_ROLE=${RESULTS_ARRAY[2]}
URL=${RESULTS_ARRAY[3]}
VPCID=${RESULTS_ARRAY[4]}

printf "${PRIMARY}* Authenticating with EC2 Container Repository${NC}\n";

#`aws ecr get-login --region $REGION --no-include-email`
#echo $(aws ecr get-login-password --region ap-southeast-1) | docker login --password-stdin --username phuongcao 074950285369.dkr.ecr.us-east-1.amazonaws.com/ecsworker
aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 074950285369.dkr.ecr.ap-southeast-1.amazonaws.com

# Tag for versioning the container images, currently set to timestamp
TAG=`date +%s`

for SERVICE_NAME in "${DEPLOYABLE_SERVICES[@]}"
do
	printf "${PRIMARY}* Locating the ECR repository for service \`${SERVICE_NAME}\`${NC}\n";

	# Find the ECR repo to push to
	printf "${PRIMARY}* Finding ECR repository for service \`${SERVICE_NAME}\`${NC}\n";
	REPO=`aws ecr describe-repositories \
		--region $REGION \
		--repository-names "$SERVICE_NAME" \
		--query "repositories[0].repositoryUri" \
		--output text`
			
	if [ "$?" != "0" ]; then
		# The repository was not found, create it
		printf "${PRIMARY}* Creating new ECR repository for service \`${SERVICE_NAME}\`${NC}\n";

		REPO=`aws ecr create-repository \
			--region $REGION \
			--repository-name "$SERVICE_NAME" \
			--query "repository.repositoryUri" \
			--output text`
		printf "${PRIMARY}* Getting ECR repository: \`${REPO}\`${NC}\n"

	fi

	printf "${PRIMARY}* Building \`${SERVICE_NAME}\`${NC}\n";

	# Build the container, and assign a tag to it for versioning
	# (cd services/$SERVICE_NAME && npm install);
	docker build -t $SERVICE_NAME -f ./src/core/$SERVICE_NAME/build/Dockerfile ./src/core/$SERVICE_NAME
	docker tag $SERVICE_NAME:latest $REPO:$TAG
	# Push the tag up so we can make a task definition for deploying it
	printf "${PRIMARY}* Pushing \`${SERVICE_NAME}\`${NC}\n";
	docker push $REPO:$TAG

	printf "${PRIMARY}* Creating new task definition for \`${SERVICE_NAME}\`${NC}\n";

	# Build an create the task definition for the container we just pushed
	CONTAINER_DEFINITIONS=$(cat <<-EOF
		[{
			"name": "$SERVICE_NAME",
			"image": "$REPO:$TAG",
			"cpu": 256,
			"memory": 256,
			"portMappings": [{
				"containerPort": 3000,
				"hostPort": 3000
		}],
		"secrets": [
			{
				"name": "API_PORT",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_API_PORT"
			},
			{
				"name": "DB_URL",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_DB_URL"
			},
			{
				"name": "AUTH0_DOMAIN",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_AUTH0_DOMAIN"
			},
			{
				"name": "AUTH0_AUDIENCE",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_AUTH0_AUDIENCE"
			},
			{
				"name": "AUTH0_MANAGEMENT_CLIENT_ID",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_AUTH0_MANAGEMENT_CLIENT_ID"
			},
			{
				"name": "AUTH0_MANAGEMENT_CLIENT_SECRET",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_AUTH0_MANAGEMENT_CLIENT_SECRET"
			},
			{
				"name": "AUTH0_CREATOR_ROLE_ID",
				"valueFrom": "arn:aws:ssm:ap-southeast-1:074950285369:parameter/devopscomvn/api/development/IDENTITY_AUTH0_CREATOR_ROLE_ID"
			}
		],
		"essential": true
		}]
	EOF
	)
	printf "${PRIMARY}* CONTAINER_DEFINITIONS: \`${CONTAINER_DEFINITIONS}\`${NC}\n";


	TASK_DEFINITION_ARN=`aws ecs register-task-definition \
		--region $REGION \
		--family $SERVICE_NAME \
		--container-definitions "$CONTAINER_DEFINITIONS" \
		--query "taskDefinition.taskDefinitionArn" \
		--output text`

	printf "${PRIMARY}* TASK_DEFINITION_ARN: \`${TASK_DEFINITION_ARN}\`${NC}\n";

	# Ensure that the service exists in ECS
	STATUS=`aws ecs describe-services \
		--region $REGION \
		--cluster $CLUSTER_NAME \
		--services $SERVICE_NAME \
		--query "services[0].status" \
		--output text`
	printf "${PRIMARY}* STATUS: \`${STATUS}\`${NC}\n";

	if [ "$STATUS" != "ACTIVE" ]; then
		# New service that needs to be deployed because it hasn't
		# been created yet.
		if [ -e "./services/$SERVICE_NAME/rule.json" ]; then
			# If this service has a rule setup for routing traffic to the service, then
			# create a target group for the service, and a rule on the ELB for routing
			# traffic to the target group.
			printf "${PRIMARY}* Setting up web facing service \`${SERVICE_NAME}\`${NC}\n";
			printf "${PRIMARY}* Creating target group for service \`${SERVICE_NAME}\`${NC}\n";

			TARGET_GROUP_ARN=`aws elbv2 create-target-group \
				--region $REGION \
				--name $SERVICE_NAME \
				--vpc-id $VPCID \
				--port 80 \
				--protocol HTTP \
				--health-check-protocol HTTP \
				--health-check-path / \
				--health-check-interval-seconds 6 \
				--health-check-timeout-seconds 5 \
				--healthy-threshold-count 2 \
				--unhealthy-threshold-count 2 \
				--query "TargetGroups[0].TargetGroupArn" \
				--output text`

			printf "${PRIMARY}* Locating load balancer listener \`${SERVICE_NAME}\`${NC}\n";

			LISTENER_ARN=`aws elbv2 describe-listeners \
				--region $REGION \
				--load-balancer-arn $ALB_ARN \
				--query "Listeners[0].ListenerArn" \
				--output text`

			if [ "$LISTENER_ARN" == "None" ]; then
				printf "${PRIMARY}* Creating listener for load balancer${NC}\n";

				LISTENER_ARN=`aws elbv2 create-listener \
					--region $REGION \
					--load-balancer-arn $ALB_ARN \
					--port 80 \
					--protocol HTTP \
					--query "Listeners[0].ListenerArn" \
					--default-actions Type=forward,TargetGroupArn=$TARGET_GROUP_ARN \
					--output text`
			fi

			printf "${PRIMARY}* Adding rule to load balancer listener \`${SERVICE_NAME}\`${NC}\n";

			# Manipulate the template to customize it with the target group and listener
			RULE_DOC=`cat ./services/$SERVICE_NAME/rule.json |
								jq ".ListenerArn=\"$LISTENER_ARN\" | .Actions[0].TargetGroupArn=\"$TARGET_GROUP_ARN\""`

			aws elbv2 create-rule \
				--region $REGION \
				--cli-input-json "$RULE_DOC"

			printf "${PRIMARY}* Creating new web facing service \`${SERVICE_NAME}\`${NC}\n";

			LOAD_BALANCERS=$(cat <<-EOF
				[{
					"targetGroupArn": "$TARGET_GROUP_ARN",
					"containerName": "$SERVICE_NAME",
					"containerPort": 3000
				}]
			EOF
			)

			RESULT=`aws ecs create-service \
				--region $REGION \
				--cluster $CLUSTER_NAME \
				--load-balancers "$LOAD_BALANCERS" \
				--service-name $SERVICE_NAME \
				--role $ECS_ROLE \
				--task-definition $TASK_DEFINITION_ARN \
				--desired-count 1`
		else
			# This service doesn't have a web interface, just create it without load balancer settings
			printf "${PRIMARY}* Creating new background service \`${SERVICE_NAME}\`${NC}\n";
			printf "CLUSTER NAME: ${CLUSTER_NAME}"
			printf "SERVICE NAME: ${SERVICE_NAME}"
			printf "TASK DEFINITION ARN: ${TASK_DEFINITION_ARN}"
			RESULT=`aws ecs create-service \
				--region $REGION \
				--cluster $CLUSTER_NAME \
				--service-name $SERVICE_NAME \
				--task-definition $TASK_DEFINITION_ARN \
				--desired-count 1`
		fi
	else
		# The service already existed, just update the existing service.
		printf "${PRIMARY}* Updating service \`${SERVICE_NAME}\` with task definition \`${TASK_DEFINITION_ARN}\`${NC}\n";
		printf "${PRIMARY}* CLUSTER_NAME \`${CLUSTER_NAME}\` with task definition \`${TASK_DEFINITION_ARN}\`${NC}\n";

		RESULT=`aws ecs update-service \
			--region $REGION \
			--cluster $CLUSTER_NAME \
			--service $SERVICE_NAME \
			--task-definition $TASK_DEFINITION_ARN`
	fi
done

printf "${PRIMARY}* Done, application is at: http://${URL}${NC}\n";
printf "${PRIMARY}* (It may take a minute for the container to register as healthy and begin receiving traffic.)${NC}\n";
