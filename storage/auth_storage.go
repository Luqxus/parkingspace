package storage

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/luquxSentinel/spacedrive/types"
)

type AuthStorage interface {
	CountEmail(ctx context.Context, email string) (int64, error)

	CreateUser(ctx context.Context, user *types.User) error

	GetUserWithEmail(ctx context.Context, email string) (*types.User, error)

	UpdateLastSignIn(ctx context.Context, email string) error
}

type authstorage struct {
	db *sql.DB
}

func NewAuthStorage() (*authstorage, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	return &authstorage{
		db: db,
	}, nil
}

func (s *authstorage) CountEmail(ctx context.Context, email string) (int64, error) {
	query := `SELECT COUNT(email) FROM Users WHERE email = $1`
	result, err := s.db.Query(query, email)
	if err != nil {
		return -1, err
	}

	var count int64
	err = result.Scan(&count)
	return count, err
}

func (s *authstorage) CreateUser(ctx context.Context, user *types.User) error {
	query := `INSERT INTO Users (uid, email, first_name, last_name, password, created_at)
		      VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Exec(query, user.UID, user.Email, user.FirstName, user.LastName, user.Password, user.CreatedAt)

	return err
}

func (s *authstorage) GetUserWithEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT * FROM Users WHERE email = $1`

	row := s.db.QueryRow(query, email)

	user := new(types.User)

	err := row.Scan(&user.UID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authstorage) UpdateLastSignIn(ctx context.Context, email string) error {
	query := `UPDATE Users
			  SET last_sign_in = $1
			  WHERE email = $2`

	_, err := s.db.Exec(query, time.Now().Local(), email)
	return err

}

func connect() (*sql.DB, error) {
	DB_URI := os.Getenv("POSTGRES_URI")
	if DB_URI == "" {
		return nil, errors.New("POSTGRES_URI not found")
	}
	db, err := sql.Open("postgres", DB_URI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, err
}
