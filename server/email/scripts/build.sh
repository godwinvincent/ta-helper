#!/usr/bin/env bash
cd ./..

GOOS=linux go build

docker build -t info441tapal/email .

# docker push info441tapal/gateway

# go clean