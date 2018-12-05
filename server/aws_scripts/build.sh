#!/usr/bin/env bash

cd ./../gateway
GOOS=linux go build
docker build -t info441tapal/gateway .
docker push info441tapal/gateway
go clean

cd ./../mongo
# lots of work, need to build docker file