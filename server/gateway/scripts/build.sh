#!/usr/bin/env bash
cd ./..

GOOS=linux go build

docker build -t pattyold/gateway .

#docker push pattyold/gateway

go clean