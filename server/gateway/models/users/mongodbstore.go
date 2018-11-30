package users

//MySQLStore represents a store
type MongoStore struct {
	// Client monogdb
}

//NewMySQLStore constructs a new MySQLStore.
func CreateMongoStore() {
	// return &MongoStore{}
}

//GetByID returns the User with the given ID
func (s *MongoStore) GetByID(id int64) (*User, error) {
	return &User{}, nil
}

//GetByEmail returns the User with the given email
func (s *MongoStore) GetByEmail(email string) (*User, error) {
	return &User{}, nil
}

//GetByUserName returns the User with the given Username
func (s *MongoStore) GetByUserName(username string) (*User, error) {
	return &User{}, nil
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (s *MongoStore) Insert(user *User) (*User, error) {
	return &User{}, nil
}

//Update applies UserUpdates to the given user ID
//and returns the newly-updated user
func (s *MongoStore) Update(id int64, updates *Updates) (*User, error) {
	return &User{}, nil
}

//Delete deletes the user with the given ID
func (s *MongoStore) Delete(id int64) error {
	return nil
}
