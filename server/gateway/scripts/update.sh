cd ./..

docker rm -f gateway

docker run -d --name gateway \
--network ta-pal \
-e ADDR=$ADDR \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-p 80:80 \
info441tapal/gateway