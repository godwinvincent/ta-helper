cd ./..

docker rm -f gateway

docker pull info441tapal/gateway

docker run -d --name gateway \
-p 80:80 \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
info441tapal/gateway