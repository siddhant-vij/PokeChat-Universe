#!/bin/bash

source .env
cd scripts/sql/schema

countFiles=$(ls -1q . | wc -l)
if [ $countFiles -gt 0 ];
then
  for count in $(seq 1 $countFiles);
  do
    goose postgres "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_DATABASE?sslmode=disable" down
  done
fi

redis-cli flushall