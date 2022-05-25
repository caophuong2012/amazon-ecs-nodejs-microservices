package models

import "time"

type User struct {
	ID        string // uuid
	Auth0ID   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
