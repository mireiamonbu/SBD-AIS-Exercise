#!/bin/sh
set -e

# Check if the container with the required names exist, if they do, delete them 
docker rm -f orderservice orders-postgres 

# The code bellow starts the PostgreSQL container and expose its dataset port (5432)
docker run -d \
  --name orders-postgres \
  -p 5432:5432 \
  -p 3000:3000 \
  --env-file debug.env \
  -e PGDATA=/var/lib/postgresql/18/docker \
  -v pg18:/var/lib/postgresql/18/docker \
  postgres:18

# Then build the servive image from the Dockerfile
docker build -t orderservice .

# Run the service using Postgres' network namespace
docker run -d \
  --name orderservice \
  --network container:orders-postgres \
  --env-file debug.env \
  -e PGSSLMODE=disable \
  orderservice

echo " Everything prepare!"