package models

import "time"

type Creator struct {
	ID        string // uuid
	UserID    string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
