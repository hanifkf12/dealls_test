package entity

import "time"

type Transaction struct {
	ID          int       `json:"id,omitempty" db:"id"`
	UserID      int       `json:"user_id,omitempty" db:"user_id"`
	Price       float64   `json:"price,omitempty" db:"price"`
	PackageType string    `json:"package_type,omitempty" db:"package_type"` // 7_day || 1_week || 1_month
	ValidUntil  time.Time `json:"valid_until" db:"valid_until"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
