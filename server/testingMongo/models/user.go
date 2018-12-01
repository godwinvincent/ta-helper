package models

type UserModel struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

func NewUserModel(username string, email string) *UserModel {
	return &UserModel{username, email}
}
