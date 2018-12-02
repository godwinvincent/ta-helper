#!/bin/bash



export ADDR=:80
export REDISADDR=localhost:6379
export SESSIONKEY=testKey
export MONGOADDR=localhost:27017
export MONGODB=ta-pal

cd ../gateway/
go clean
go build
./gateway
cd ../testingMongo