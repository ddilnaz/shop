//pkg\abr-plus\model\user_model.go
package model

import (
	"database/sql"
	"errors"
	"log"
)
var users = []User{
	{Id: 1, Name: "Zeinolla Dilnaz", Username: "zeinolla_d", Password: "password123"},
	{Id: 2, Name: "Zhumanova Zhanel", Username: "ahanel_zh", Password: "pass456"},
	{Id: 3, Name: "Ali John", Username: "ali_zh", Password: "secret123"},
	{Id: 4, Name: "Kami White", Username: "kami_w", Password: "exam_pass"},
	{Id: 5, Name: "Airat Green", Username: "airat_g", Password: "green"},
}

type UserModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}


func (um *UserModel) GetAllUsers() ([]User, error) {
	return users, nil
}

func (um *UserModel) GetUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.Id == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (um *UserModel) CreateUser(user *User) error {
	// Simulating auto-increment ID
	user.Id = len(users) + 1
	users = append(users, *user)
	return nil
}

func (um *UserModel) UpdateUser(user *User) error {
	for i, u := range users {
		if u.Id == user.Id {
			users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (um *UserModel) DeleteUser(id int) error {
	for i, user := range users {
		if user.Id == id {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
