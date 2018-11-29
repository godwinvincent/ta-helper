package mongo

import (
	"gopkg.in/mgo.v2"
)

type MongoSession struct {
	session *mgo.Session
}

func NewSession(url string) (*MongoSession, error) {
	// Use the URL to make a connection to that URL
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoSession{session}, err
}

func (s *MongoSession) Copy() *MongoSession {
	return &MongoSession{s.session.Copy()}
}

func (s *MongoSession) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *MongoSession) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
