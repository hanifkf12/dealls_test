package entity

import "time"

type Profile struct {
	ID        int        `json:"id,omitempty" db:"id"`
	UserID    int        `json:"user_id,omitempty" db:"user_id"`
	Name      string     `json:"name,omitempty" db:"name"`
	Avatar    string     `json:"avatar,omitempty" db:"avatar"`
	Age       int        `json:"age,omitempty" db:"age"`
	Gender    string     `json:"gender,omitempty" db:"gender"`
	Bio       string     `json:"bio,omitempty" db:"bio"`
	Location  string     `json:"location,omitempty" db:"location"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
