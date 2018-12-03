# Questions


**Change multiple orders**

Yes, mongo allows for updating multiple documents. Just use a modifier operation and multi=True. For example, this increments order by one for all documents with order greater than five:


Where condition on an update: [stack overflow](https://stackoverflow.com/questions/13420720/ordering-for-a-todo-list-with-mongodb)

`todos.update({'order':{'$gt':5}}, {'$inc':{'order':1}}, multi=True)`

```
    c.collections.updateAll({ "$and": []bson.M{ bson.M{"officeHoursID":"____"},'questionOrder':{'$gt':5}  } }, {'$inc':{'order':1}}, multi=True)


    Example
    err = c.Find( bson.M{
         "$or": []bson.M{
              bson.M{"uuid":"UUID0"}, 
              bson.M{"name": "Joe"} 
              } 
         } ).One(&result)
```
Find highest order


****



