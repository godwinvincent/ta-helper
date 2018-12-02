#!/usr/bin/env bash
cd ./..

GOOS=linux go build

docker build -t info441tapal/gateway .

# docker push info441tapal/gateway

# go clean