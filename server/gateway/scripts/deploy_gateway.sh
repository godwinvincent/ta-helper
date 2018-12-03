#!/usr/bin/env bash




# ---- Build Gateway ----
cd ./..
GOOS=linux go build
docker build -t info441tapal/gateway .
# docker push info441tapal/gateway
go clean

cd scripts


# ---- Deploy onto AWS ----
# ssh -i "TA-Pal-API.pem" ec2-user@ec2-18-188-13-55.us-east-2.compute.amazonaws.com < gateway_helper.sh

# ---- Deploy Locally ----
./gateway_helper.sh