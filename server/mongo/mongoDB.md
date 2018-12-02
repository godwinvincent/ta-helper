# Mongo DB

How to connect to Mongo DB if it's in a container
```
docker run -it \
--rm \
--network ta-pal \
mongo sh -c 'exec mongo mongo:27017'

```



## Mongo

**MGO**
- this is a golang to mongo interface/package
- Golang to Mongo Interface is MGO: [code examples](https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e).


**Mongo Shell commands**
- use tahelper
- show collections
    - db.createCollection("users")
    - db.createCollection("questions")
    - db.<collection_name>.drop()
- Mongo [CLI commands](https://dzone.com/articles/top-10-most-common-commands-for-beginners).


