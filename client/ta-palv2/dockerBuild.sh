#!/bin/bash
docker pull godwinvincent/info441-client
docker rm -f 441client
docker run -d --name 441client -p 443:443 -p 80:80 -v /etc/letsencrypt:/etc/letsencrypt:ro godwinvincent/info441-client