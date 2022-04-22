package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ZweWT/Go-Test.git/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

var ErrDuplicateEmail = errors.New("duplicate email")

//plain text field is a pointer to string to be able to distinguish the plaintext not being present at all in struct, or empty strinf
type password struct {
	plaintext *string
	hash      []byte
}

type UserModel struct {
	DB *sql.DB
}

func (p *password) Set(plaintextPw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPw), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPw
	p.hash = hash

	return nil
}

func (p *password) Match(plaintextPw string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPw))
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

func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users(name, email, password_hash)
		VALUES ($1, $2, $3)`

	args := []interface{}{user.Name, user.Email, user.Password.hash}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail

		default:
			return err
		}
	}
	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	query := `
	SELECT id, name, email, password_hash, created_at 
	FROM users
	WHERE email = $1`
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
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

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "is required")

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (m UserModel) Get(id int64) (*User, error) {
	query := `
	SELECT id, created_at, name, email, password_hash
	FROM users
	WHERE id = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.hash,
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
