docker rm -f client
docker pull info441tapal/client

docker run -d --name client \
info441tapal/client