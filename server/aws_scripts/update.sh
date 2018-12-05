docker rm -f gateway
docker rm -f mongo
docker rm -f redis
docker rm -f schedule
docker rm -f rabbit

docker pull info441tapal/gateway
docker pull info441tapal/schedule

docker run -d --name gateway \
--network ta-pal \
-e ADDR=$ADDR \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-p 80:80 \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
info441tapal/gateway

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