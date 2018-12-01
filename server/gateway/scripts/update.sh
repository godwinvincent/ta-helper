cd ./..

docker rm -f gateway

docker run -d --name gateway \
-p 80:80 \
info441tapal/gateway