//C:\Users\Lenovo\Desktop\shop\pkg\abr-plus\model\user_model.go
package model
import (
	"database/sql"
	"log"
	"context"
	"time"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/ddilnaz/shop/pkg/shop_tour/validator"
	"crypto/sha256"
)
type User struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Activated bool      `json:"activated"`
	Password  password `json:"-"`
	Version   int       `json:"-"`
}

var users = []User{
	{Id: 1, Name: "Zeinolla Dilnaz", Email: "zeinolla_d@gmail.com"},
	{Id: 2, Name: "Zhumanova Zhanel", Email: "ahanel_zh"},
	{Id: 3, Name: "Ali John", Email: "ali_zh"},
	{Id: 4, Name: "Kami White", Email: "kami_w"},
	{Id: 5, Name: "Airat Green", Email: "airat_g"},
}
var AnonymousUser = &User{} 
var (
	ErrDuplicateEmail = errors.New("duplicate email")
)
type UserModel struct{
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
type password struct {
	plaintext *string
	hash      []byte
}
func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash
	return nil
}
func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
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
		INSERT INTO users (name, email, password_hash, activated)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version
		`

	args := []interface{}{user.Name, user.Email, user.Password.hash, user.Activated}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pqErr := `pq: duplicate key value violates unique constraint "users_email_key"`
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == pqErr:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}
func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
		SELECT id, created_at, name, email, password_hash, activated, version
		FROM users
		WHERE email = $1
		`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

// Update updates the details for a specific user in the users table. Note, we check against the
// version field to hel
func (m UserModel) UpdateUser(user *User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password_hash = $3, activated = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version
		`

	args := []interface{}{
		user.Name,
		user.Email,
		user.Password.hash,
		user.Activated,
		user.Id,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// GetForToken retrieves a user record from the users table for an associated token and token scope.
func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
	// Calculate the SHA-256 hash for the plaintext token provided by the client.
	// Note, that this will return a byte *array* with length 32, not a slice.
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
		SELECT 
			users.id, users.created_at, users.name, users.email, 
			users.password_hash, users.activated, users.version
		FROM       users
        INNER JOIN tokens
			ON users.id = tokens.user_id
        WHERE tokens.hash = $1  --<-- Note: this is potentially vulnerable to a timing attack, 
            -- but if successful the attacker would only be able to retrieve a *hashed* token 
            -- which would still require a brute-force attack to find the 26 character string
            -- that has the same SHA-256 hash that was found from our database. 
			AND tokens.scope = $2
			AND tokens.expiry > $3
		`

	// Create a slice containing the query args. Note, that we use the [:] operator to get a slice
	// containing the token hash, since the pq driver does not support passing in an array.
	// Also, we pass the current time as the value to check against the token expiry.
	args := []interface{}{tokenHash[:], tokenScope, time.Now()}

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Execute the query, scanning the return values into a User struct. If no matching record
	// is found we return an ErrRecordNotFound error.
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// Return the matching user.
	return &user, nil
}

// ValidateEmail checks that the Email field is not an empty string and that it matches the regex
// for email addresses, validator.EmailRX.
func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be valid email address")
}

// ValidatePasswordPlaintext validtes that the password is not an empty string and is between 8 and
// 72 bytes long.
func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	// validate user.Name
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	// Validate email
	ValidateEmail(v, user.Email)

	// If the plaintext password is not nil, call the standalone ValidatePasswordPlaintext helper.
	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	// If the password has is ever nil, this will be due to a logic error in our codebase
	// (probably because we forgot to set a password for the user). It's a useful sanity check to
	// include here, but it's not a problem with the data provided by the client. So, rather
	// than adding an error to the validation map we raise a panic instead.
	if user.Password.hash == nil {
		// TODO: fix this panic
		panic("missing password hash for user")
	}
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
