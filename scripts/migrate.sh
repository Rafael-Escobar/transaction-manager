#!/bin/bash

if [[ -z $1 ]]; then
  echo "Please select the enviroment: stg, or prd"
  exit 1
fi

if [[ -z $2 ]]; then
  echo "Please select the action: up, down"
  exit 1
fi

ENV=$1
DB_TYPE="postgres"
if [[ "$ENV" == "stg" ]] || [[ "$ENV" == "prd" ]]
then
  DB_USER="user" ## Change this to get from remote secret manager
  DB_PASSWORD="password" ## Change this to get from remote secret manager
  DB_NAME="transaction-manager-db-${ENV}"
else
  DB_USER="user"
  DB_PASSWORD="password"
  DB_NAME="transaction-manager-db-${ENV}"
fi

if [[ $2 == "up" ]]; then
  migrate -path './migrations' -database "$DB_TYPE://$DB_USER:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable" up
fi

if [[ $2 == "down" ]]; then
  migrate -path './migrations' -database "$DB_TYPE://$DB_USER:$DB_PASSWORD@localhost:5432/$DB_NAME?sslmode=disable" down 1
fi
