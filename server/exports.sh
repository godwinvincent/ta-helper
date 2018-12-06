#!/bin/sh

export ADDR=:80
export REDISADDR=redis:6379
export SESSIONKEY=testKey
export MONGOADDR=mongo:27017
export MONGODB=ta-pal
export TLSCERT=/etc/letsencrypt/live/tapalapi.patrickold.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/tapalapi.patrickold.me/privkey.pem
# export RABBITADDR=localhost:5672
