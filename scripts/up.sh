#!/bin/bash

source .env
cd scripts/sql/schema
goose postgres "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable" up