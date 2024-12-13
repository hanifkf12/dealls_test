package entity

import "time"

type Profile struct {
	ID        int
	UserID    int
	Name      string
	Avatar    string
	Age       int
	Gender    string
	Bio       string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
