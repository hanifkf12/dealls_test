package entity

import "time"

type Swipe struct {
	ID        int       `json:"id,omitempty"`
	UserID    int       `json:"user_id,omitempty"`
	ProfileID int       `json:"profile_id,omitempty"`
	SwipeType int       `json:"swipe_type,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
