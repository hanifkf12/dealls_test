package entity

import "time"

type Transaction struct {
	ID        int
	UserID    int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
