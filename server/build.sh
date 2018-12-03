#!/usr/bin/env bash

source exports.sh

cd ./scheduling


# ----- Build -----
GOOS=linux go build
docker build -t info441tapal/schedule .
go clean


# ----- Deploy -----
docker rm -f schedule

docker run -d --name schedule \
--network ta-pal \
-e ADDR=$ADDR \
-e REDISADDR=$REDISADDR \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
info441tapal/schedule