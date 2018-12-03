package models

import mgo "gopkg.in/mgo.v2"

type MongoSession struct {
	Session *mgo.Session
}

type MongoCollection struct {
	Collection *mgo.Collection
}
