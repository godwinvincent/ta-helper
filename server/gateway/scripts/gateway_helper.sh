#!/usr/bin/env bash


docker rm -f gateway

# docker pull info441tapal/gateway

docker run -d --name gateway \
--network ta-pal \
-e ADDR=$ADDR \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-p 80:80 \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
info441tapal/gateway