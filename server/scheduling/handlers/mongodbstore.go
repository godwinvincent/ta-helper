package handlers

/**
 * Hi, welcome to Ben's mongo interface code.
 * If you want to work with a mongo DB you need to:
 * 1) make a db connection
 * 2) make a collection struc
 * 3) use that collection struc to call a function.
 */

// ------------- Code from Gateway -------------

//User represents a user account in the database
type User struct {
	Email          string `json:"email" bson:"email"`
	PassHash       []byte `json:"-" bson:"passHash"` //never JSON encoded/decoded
	UserName       string `json:"username" bson:"username"`
	FirstName      string `json:"firstName" bson:"firstName"`
	LastName       string `json:"lastName" bson:"lastName"`
	Role           string `json:"role" bson:"role"`
	EmailActivated bool   `json:"emailActivated" bson:"emailActivated"`
	EmailVerifCode string `json:"-" bson:"emailVerifCode"`
}

// // GetByUserName retrives a user from the given collection and returns it as a User
// func (col *models.MongoCollection) GetByUserName(username string) (*User, error) {
// 	model := User{}
// 	err := col.Collection.Find(bson.M{"username": username}).One(&model)
// 	return &model, err
// }
