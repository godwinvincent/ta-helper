#!/bin/sh

# -------- Push Mongo DB -------

# docker build -t bwalchen/gateway .
# docker push bwalchen/gateway

# ssh -i /Users/ben/.ssh/i441_key.pem ec2-user@441api.walchen.com < mongo_helper.sh

docker rm -f rabbit
docker rm -f mongo
docker rm -f redis

# Run RabbitMQ
docker run -d --hostname my-rabbit \
--name rabbit \
--network ta-pal \
rabbitmq:3

# Run Mongo
docker run -d \
-p 27017:27017 \
--name mongo \
--network ta-pal \
mongo


# Run Redis
docker run -d --name redis \
--network ta-pal \
-p 6379:6379 \
redis


