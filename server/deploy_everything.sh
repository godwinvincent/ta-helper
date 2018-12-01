#!/bin/sh

docker rm -f gateway
docker rm -f mongo
docker rm -f redis


docker network remove ta-pal
docker network create ta-pal

cd mongo 
./deploy_mongo.sh

cd ../gateway/scripts
deploy.sh