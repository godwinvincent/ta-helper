#!/usr/bin/env bash
GOOS=linux go build
docker build -t info441tapal/schedule .
go clean
docker rm -f schedule
docker run -d --name schedule \
--network ta-pal \
-e ADDR=$ADDR \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
info441tapal/schedule