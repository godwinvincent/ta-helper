# Questions


**Change multiple orders**

Yes, mongo allows for updating multiple documents. Just use a modifier operation and multi=True. For example, this increments order by one for all documents with order greater than five:

`todos.update({'order':{'$gt':5}}, {'$inc':{'order':1}}, multi=True)`




Find highest order


****



https://stackoverflow.com/questions/13420720/ordering-for-a-todo-list-with-mongodb