package types

import "time"

type User struct {
	UID        string `db:"uid"`
	Email     string	`db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	LastSignIn time.Time `db:"last_sign_in"`
	IsEmailValified bool `db:"is_email_verified"`
	CreatedAt time.Time	`db:"created_at"`
}

type CreateUserData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
