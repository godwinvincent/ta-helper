#!/bin/sh

# !!! Import critical environment variables and creds
source exports.sh

docker rm -f gateway
docker rm -f mongo
docker rm -f redis
docker rm -f schedule
docker rm -f rabbit


docker network remove ta-pal
docker network create ta-pal

# Deploy RabbitMQ, Mongo, and Redis
cd mongo 
./deploy_mongo.sh
cd ..

# Give time for everything to build before 
# having Gateway connect to them
sleep 10

# Deploy Gateway
cd gateway/scripts/
./deploy_gateway.sh