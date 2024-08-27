#!/bin/bash

source .env

pg_dump -U $DB_USERNAME $DB_DATABASE > pgDump_$DB_DATABASE.sql