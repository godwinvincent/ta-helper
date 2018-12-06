cd ./..

docker rm -f email

docker run -d --name email \
--network ta-pal \
-e ADDR=$ADDR \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e MONGOADDR=$MONGOADDR \
-e MONGODB=$MONGODB \
-e SENDGRID_API_KEY=$SENDGRID_API_KEY \
info441tapal/email