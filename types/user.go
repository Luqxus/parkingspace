package types

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
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
