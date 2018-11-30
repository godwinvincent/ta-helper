package users

import (
	"errors"
	"time"

	"github.com/godwinvincent/homework-godwinvincent/servers/gateway/indexes"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//Store represents a store for Users
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *Updates) (*User, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	LogSuccesfulLogIn(id int64, time time.Time, ipAddress string) error

	//Delete deletes the user with the given ID
	Delete(id int64) error

	//Enrolls user in 2FA
	EnrollIn2FA(id int64, twoFaEnabled bool, twoFaStruct []byte) error

	//Adds all Users to Trie
	AddAllToTrie(trie *indexes.Trie) error
}
