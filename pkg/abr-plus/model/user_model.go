//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\user_model.go
package model
import (
	"database/sql"
	"log"
	"context"
	"time"
	
)
type User struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var users = []User{
	{Id: 1, Name: "Zeinolla Dilnaz", Email: "zeinolla_d@gmail.com"},
	{Id: 2, Name: "Zhumanova Zhanel", Email: "ahanel_zh"},
	{Id: 3, Name: "Ali John", Email: "ali_zh"},
	{Id: 4, Name: "Kami White", Email: "kami_w"},
	{Id: 5, Name: "Airat Green", Email: "airat_g"},
}

type UserModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}


func GetAllUsers() []User {
	return users
}


func (m UserModel) GetUserById(id int) (*User, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, name , email
		FROM users
		WHERE id = $1
		`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m UserModel) CreateUser(user *User) error {
	// Insert a new user into the database.
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, created_at`

	args := []interface{}{user.Name, user.Email}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.CreatedAt)
}

func (m UserModel) UpdateUser(user *User) error {
	// Update a specific menu item in the database.
	query := `
	UPDATE users
	SET name = $1, email = $2, created_at = $3
	WHERE id = $4
	RETURNING updated_at`

	args := []interface{}{user.Name, user.Email, user.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&user.UpdatedAt)
}


func (m UserModel) DeleteUser(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM users
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
