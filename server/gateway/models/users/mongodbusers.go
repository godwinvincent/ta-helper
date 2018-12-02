package users

/**
 * Hi, welcome to Ben's mongo interface code.
 * If you want to work with a mongo DB you need to:
 * 1) make a db connection
 * 2) make a collection struc
 * 3) use that collection struc to call a function.
 */

import (
	"github.com/alabama/final-project-alabama/server/gateway/models"
	"gopkg.in/mgo.v2/bson"
)

// -------------  Strucs -------------
type UserCollection models.MongoCollection

// ------------- Collections -------------

// InsertUser inserts a User into the given Collection
func (col *UserCollection) Insert(user *User) (*User, error) {

	if err := col.Collection.Insert(user); err != nil {
		return nil, err
	}

	newUser, err := col.GetByUserName(user.UserName)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetByUserName retrives a user from the given collection and returns it as a User
func (col *UserCollection) GetByUserName(username string) (*User, error) {
	model := User{}
	err := col.Collection.Find(bson.M{"username": username}).One(&model)
	return &model, err
}

func (col *UserCollection) GetByEmail(email string) (*User, error) {
	model := User{}
	err := col.Collection.Find(bson.M{"email": email}).One(&model)
	return &model, err
}

func (col *UserCollection) GetByID(id string) (*User, error) {
	model := User{}
	err := col.Collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&model)
	return &model, err
}

// Delete
func (col *UserCollection) Delete(id string) error {
	err := col.Collection.Remove(bson.M{"_id": id})
	return err
}

// func (col *MongoCollection) Update(id string, newUser *User) error {

// 	err := col.collection.UpdateId(obj1.Id, bson.M{"$set": &obj1})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
