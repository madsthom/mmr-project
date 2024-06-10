#!/bin/bash

# Exit on error
set -e

# Remove the existing dump - may still be there if the script failed
rm -rf postgres_dump

# Variables

serverName=${1:-$DB_SERVER_NAME}
resourceGroupName=${2:-$DB_RESOURCE_GROUP}
databaseName=${3:-$DB_NAME}
tenantId=${4:-$AZURE_TENANT_ID}
subscriptionId=${5:-$AZURE_SUBSCRIPTION_ID}
username=${5:-$DB_USERNAME}

# Login to Azure
az login --tenant $tenantId

# Set the subscription
az account set --subscription $subscriptionId

# Get the fully qualified server name
serverFQDN=$(az postgres flexible-server show --name $serverName --resource-group $resourceGroupName --query fullyQualifiedDomainName -o tsv)

# Get an access token
accessToken=$(az account get-access-token --resource-type oss-rdbms --query accessToken -o tsv)

# Connect to the database
PGPASSWORD=$accessToken pg_dump -Fd -n public $databaseName -h $serverFQDN -p 5432 -U $username -f postgres_dump

# Drop the database
PGPASSWORD=this_is_a_hard_password1337 psql -d $databaseName -h localhost -p 5432 -U postgres -c "drop schema if exists public cascade;"

# Import the data
PGPASSWORD=this_is_a_hard_password1337 pg_restore -Fd -d $databaseName postgres_dump -h localhost -p 5432 -U postgres

# Remove the dump
rm -rf postgres_dump
