#!/usr/bin/env bash

# Write RDS demo parameters to Parameter Store

# Put parameters into Parameter Store
aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_HOST \
  --type String \
  --value "172.0.0.1" \
  --description "DB Host" \
  --overwrite

aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_NAME \
  --type String \
  --value "db_identity" \
  --description "DB Name" \
  --overwrite

aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_PORT \
  --type String \
  --value "5432" \
  --description "DB Port" \
  --overwrite

aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_USER \
  --type String \
  --value "identity" \
  --description "DB User" \
  --overwrite

aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_PASS \
  --type String \
  --value "password" \
  --description "DB Password" \
  --overwrite

aws ssm put-parameter \
  --name /awst/api/development/API_IDENTITY_DB_MAX_OPEN_CONNS \
  --type String \
  --value "100" \
  --description "DB Maxe Open Connections" \
  --overwrite

# Get parameters from Parameter Store
aws ssm get-parameter \
  --name /awst/api/development/API_IDENTITY_DB_HOST \
  #--query Parameter.Value