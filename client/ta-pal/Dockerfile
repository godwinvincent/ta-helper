# base image
FROM node:latest

# set working directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# install and cache app dependencies
COPY package.json /usr/src/app/
RUN npm install

ADD src /usr/src/app/src
ADD public /usr/src/app/public

RUN npm build

# start app
CMD ["npm", "start"]