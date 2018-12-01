package users

import (
	"crypto/md5"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//gravatarBasePhotoURL is the base URL for Gravatar image requests.
//See https://id.gravatar.com/site/implement/images/ for details
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

//bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

//User represents a user account in the database
type User struct {
	Email     string `json:"email" bson:"email"`
	PassHash  []byte `json:"-" bson:"passHash"` //never JSON encoded/decoded
	UserName  string `json:"username" bson:"username"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`

	EmailActivated bool   `json:"-" bson:"emailActivated"` //never JSON encoded/decoded
	EmailVerifCode string `json:"-" bson:"emailVerifCode"`
}

//Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

//Updates represents allowed updates to a user profile
type Updates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	if _, err := mail.ParseAddress(nu.Email); err != nil {
		return fmt.Errorf("Invalid Email Address")
	}
	if len(nu.Password) < 6 {
		return fmt.Errorf("Invalid Password")
	}
	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("Password and Confirmation don't match")
	}
	if len(nu.UserName) == 0 || strings.ContainsAny(nu.UserName, " ") {
		return fmt.Errorf("Invalid Username")
	}
	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	if err := nu.Validate(); err != nil {
		return nil, err
	}
	user := User{}
	user.FirstName = nu.FirstName
	user.LastName = nu.LastName
	user.Email = nu.Email
	user.UserName = nu.UserName
	text := []byte(strings.ToLower(strings.TrimSpace(user.Email)))
	hasher := md5.New()
	hasher.Write(text)

	user.SetPassword(nu.Password)

	return &user, nil
}

//FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put between the names. If both are missing,
//this returns an empty string
func (u *User) FullName() string {
	if u.FirstName != "" && u.LastName != "" {
		return u.FirstName + " " + u.LastName
	} else if u.FirstName != "" {
		return u.FirstName
	} else if u.LastName != "" {
		return u.LastName
	}
	return ""
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil
	}
	u.PassHash = hash
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	if updates.FirstName == "" && updates.LastName == "" {
		return fmt.Errorf("First Name and Last Name is empty")
	}
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	return nil
}
