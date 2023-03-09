package domain

import "time"

type User struct {
	Id           string
	Login        string
	Email        string
	PasswordHash string
	Confirmed    bool
	Active       bool
	DateCreated  time.Time
}
