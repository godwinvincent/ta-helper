package users

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	//For mysql connection
	_ "github.com/go-sql-driver/mysql"
	"github.com/godwinvincent/homework-godwinvincent/servers/gateway/indexes"
)

const sqlSelectAll = "SELECT id, email, username, pass_hash, first_name, last_name, photo_url, twoFaEnabled, twoFaStruct from Users "

//MySQLStore represents a store
type MySQLStore struct {
	Client *sql.DB
}

//NewMySQLStore constructs a new MySQLStore.
func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{db}
}

//GetByID returns the User with the given ID
func (s *MySQLStore) GetByID(id int64) (*User, error) {
	user := User{}
	insq := sqlSelectAll + "where id = ?"
	err := s.Client.QueryRow(insq, id).Scan(&user.ID, &user.Email, &user.UserName, &user.PassHash, &user.FirstName, &user.LastName, &user.PhotoURL, &user.TwoFA, &user.TwoFaStruct)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

//GetByEmail returns the User with the given email
func (s *MySQLStore) GetByEmail(email string) (*User, error) {
	user := User{}
	insq := sqlSelectAll + "where email = ?"
	err := s.Client.QueryRow(insq, email).Scan(&user.ID, &user.Email, &user.UserName, &user.PassHash, &user.FirstName, &user.LastName, &user.PhotoURL, &user.TwoFA, &user.TwoFaStruct)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

//GetByUserName returns the User with the given Username
func (s *MySQLStore) GetByUserName(username string) (*User, error) {
	user := User{}
	insq := sqlSelectAll + "where username = ?"
	err := s.Client.QueryRow(insq, username).Scan(&user.ID, &user.Email, &user.UserName, &user.PassHash, &user.FirstName, &user.LastName, &user.PhotoURL, &user.TwoFA, &user.TwoFaStruct)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (s *MySQLStore) Insert(user *User) (*User, error) {
	insq := "insert into Users(email, username, pass_hash, first_name, last_name, photo_url, twoFaEnabled, twoFaStruct) values (?,?,?,?,?,?,?,?)"
	res, err := s.Client.Exec(insq, user.Email, user.UserName, user.PassHash, user.FirstName, user.LastName, user.PhotoURL, user.TwoFA, user.TwoFaStruct)
	if err != nil {
		return nil, fmt.Errorf("error inserting new row: %v", err)
	}
	//get the auto-assigned ID for the new row
	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting id: %v", err)
	}
	user.ID = id
	return user, nil
}

//Update applies UserUpdates to the given user ID
//and returns the newly-updated user
func (s *MySQLStore) Update(id int64, updates *Updates) (*User, error) {
	insq := "update Users set first_name = ?, last_name = ? where id = ?"
	_, err := s.Client.Exec(insq, updates.FirstName, updates.LastName, id)
	if err != nil {
		return nil, fmt.Errorf("error altering id: %v", err)
	}
	updatedUser, err := s.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting updated user with id %d: %v", id, err)
	}

	return updatedUser, nil
}

//Delete deletes the user with the given ID
func (s *MySQLStore) Delete(id int64) error {
	_, err := s.Client.Exec("DELETE FROM users where id = ?", id) // OK
	if err != nil {
		return fmt.Errorf("error deleting id: %v", err)
	}
	return nil
}

//LogSuccesfulLogIn logs login information
func (s *MySQLStore) LogSuccesfulLogIn(id int64, time time.Time, ipAddress string) error {
	insq := "insert into UserSignIns(user_id, timestamp, ip_address) values (?,?,?)"
	_, err := s.Client.Exec(insq, id, time, ipAddress)
	if err != nil {
		return fmt.Errorf("error inserting new row: %v \t %s", err, ipAddress)
	}
	return nil
}

//EnrollIn2FA Enrolls a user in 2FA
func (s *MySQLStore) EnrollIn2FA(id int64, twoFaEnabled bool, twoFaStruct []byte) error {
	insq := "update Users set twoFaEnabled = ?, twoFaStruct = ? where id = ?"
	_, err := s.Client.Exec(insq, twoFaEnabled, twoFaStruct, id)
	if err != nil {
		return fmt.Errorf("error altering id: %v", err)
	}

	return nil
}

func (s *MySQLStore) AddAllToTrie(trie *indexes.Trie) error {
	rows, err := s.Client.Query("SELECT id, first_name, last_name, username FROM Users")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var firstName string
		var lastName string
		var username string
		err = rows.Scan(&id, &firstName, &lastName, &username)
		if err != nil {
			return err
		}
		names := strings.Split(firstName, " ")
		names = append(names, strings.Split(lastName, " ")...)
		for _, name := range names {
			trie.Add(name, id)
		}
		trie.Add(username, id)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return (err)
	}
	return nil
}
