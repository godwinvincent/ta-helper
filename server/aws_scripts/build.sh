#!/usr/bin/env bash

cd ./../gateway
GOOS=linux go build
docker build -t info441tapal/gateway .
docker push info441tapal/gateway
go clean

cd ./../scheduling
GOOS=linux go build
docker build -t info441tapal/schedule .
docker push info441tapal/schedule
go clean

cd ./../email
GOOS=linux go build
docker build -t info441tapal/email .
docker push info441tapal/email
go clean
