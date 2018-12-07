export TLSCERT=/etc/letsencrypt/live/tapal.patrickold.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/tapal.patrickold.me/privkey.pem

docker rm -f client
docker pull info441tapal/client

docker run -d --name client \
-p 80:80 \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
info441tapal/client