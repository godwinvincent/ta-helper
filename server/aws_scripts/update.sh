export ADDR=:80
export REDISADDR=redis:6379
export SESSIONKEY=testKey
export MONGOADDR=mongo:27017
export MONGODB=ta-pal
export TLSCERT=/etc/letsencrypt/live/tapalapi.patrickold.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/tapalapi.patrickold.me/privkey.pem
export RABBITADDR=rabbit:5672

docker rm -f gateway
docker rm -f mongo
docker rm -f redis
docker rm -f schedule
docker rm -f rabbit

docker pull info441tapal/gateway
docker pull info441tapal/schedule

docker network rm ta-pal
docker network create ta-pal

# Run RabbitMQ
docker run -d --hostname my-rabbit \
--name rabbit \
--network ta-pal \
rabbitmq:3

# Run Mongo
docker run -d \
--name mongo \
--network ta-pal \
-v /home/ec2-user/mongoData:/data/db \ # delete if if Mongo not working
mongo

docker run -p 27017:27017 -v /home/<user>/data:/data/db 

# Run Redis
docker run -d --name redis \
--network ta-pal \
redis

sleep 30

# Run scheduling microservice
docker run -d --name schedule \
--network ta-pal \
-e REDISADDR=$REDISADDR \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-e RABBITADDR=$RABBITADDR \
info441tapal/schedule

# Run API gateway
docker run -d --name gateway \
--network ta-pal \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-p 80:80 \
-p 443:443 \
info441tapal/gateway