#!/bin/bash
./build.sh
docker push godwinvincent/info441-client
ssh -i ~/Documents/UW/2018-2019/Autumn/INFO441/key/info-441-uw.pem ec2-user@ec2-34-215-212-175.us-west-2.compute.amazonaws.com < dockerBuild.sh