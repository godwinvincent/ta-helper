docker rm -f client
docker pull info441tapal/client

docker run -d --name client \
-p 80:80 \
-p 443:443 \
info441tapal/client