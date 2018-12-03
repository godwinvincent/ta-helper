#!/bin/sh


source exports.sh

docker rm -f gateway
docker rm -f mongo
docker rm -f redis
docker rm -f schedule


docker network remove ta-pal
docker network create ta-pal

cd mongo 
./deploy_mongo.sh
cd ..

cd gateway/scripts/
./deploy_gateway.sh