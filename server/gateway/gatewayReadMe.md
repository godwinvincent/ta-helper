# Gateway & Mongo docs


## Mongo
- Golang to Mongo Interface is MGO: [code examples](https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e).
- Mongo [CLI commands](https://dzone.com/articles/top-10-most-common-commands-for-beginners).


**DB**

How to connect to Mongo DB once it's been deployed.

```
docker run -it \
--rm \
--network ta-pal \
mongo sh -c 'exec mongo mongo:27017'
```

**Basic shell commands**
- `show dbs` 
- `db.createCollection("ta-pal");`
- `use ta-pal`
- `show collections`
    - `db.createCollection("users")`
    - `db.createCollection("questions")`
    - `db.<collection_name>.drop()`




## Required Variables
Can be found in `/server/exports.sh`

These are the require Environment variables that the API Gateway will need in order to run.
- export ADDR=:80
- export REDISADDR=localhost:6379
- export SESSIONKEY=testKey
- export MONGOADDR=localhost:27017
- export MONGODB=tahelper

