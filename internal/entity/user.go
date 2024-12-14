package entity

import "time"

type User struct {
	ID        int        `json:"id,omitempty" db:"id"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
