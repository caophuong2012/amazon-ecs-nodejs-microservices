#!/usr/bin/env bash
# Put parameters into Parameter Store
aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_API_PORT \
  --type String \
  --value "3000" \
  --description "DB Host" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_DB_URL \
  --type String \
  --value "postgres://postgres:ijSEnURsCCGfE2K@api-identity-db.cutawwotu1jb.ap-southeast-1.rds.amazonaws.com:5432/APIIdentityDBSecurityzdevelopment?sslmode=disable" \
  --description "DB_URL" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_MAX_OPEN_CONNS \
  --type String \
  --value "MAX_OPEN_CONNS" \
  --description "MAX_OPEN_CONNS" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_AUTH0_DOMAIN \
  --type String \
  --value "AUTH0_DOMAIN" \
  --description "AUTH0_DOMAIN" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_AUTH0_AUDIENCE \
  --type String \
  --value "AUTH0_AUDIENCE" \
  --description "AUTH0_AUDIENCE" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_AUTH0_MANAGEMENT_CLIENT_ID \
  --type String \
  --value "AUTH0_MANAGEMENT_CLIENT_ID" \
  --description "AUTH0_MANAGEMENT_CLIENT_ID" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_AUTH0_MANAGEMENT_CLIENT_SECRET \
  --type String \
  --value "AUTH0_MANAGEMENT_CLIENT_SECRET" \
  --description "AUTH0_MANAGEMENT_CLIENT_SECRET" \
  --overwrite

aws ssm put-parameter \
  --name /devopscomvn/api/development/IDENTITY_AUTH0_CREATOR_ROLE_ID \
  --type String \
  --value "AUTH0_CREATOR_ROLE_ID" \
  --description "AUTH0_CREATOR_ROLE_ID" \
  --overwrite

# Get parameters from Parameter Store
aws ssm get-parameter \
  --name /devopscomvn/api/development/IDENTITY_API_PORT