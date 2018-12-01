# Commands

_Below are the commands you'll need in order to get things running locally_




## Mongo
- Golang to Mongo Interface is MGO: [code examples](https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e).
- Mongo [CLI commands](https://dzone.com/articles/top-10-most-common-commands-for-beginners).

**Mongo Shell commands**
- use tahelper
- show collections
    - db.createCollection("users")
    - db.createCollection("questions")
    - db.<collection_name>.drop()



## Required Variables
These are the require Environment variables that the API Gateway will need in order to run.
- export ADDR=:80
- export REDISADDR=localhost:6379
- export SESSIONKEY=testKey
- export MONGOADDR=localhost:27017
- export MONGODB=tahelper


