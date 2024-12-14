package swipe

import (
	"context"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/pkg/databasex"
)

type swipeRepository struct {
	db databasex.Database
}

func (s *swipeRepository) SwipeRight(ctx context.Context, data entity.Swipe) (int64, error) {
	query := `INSERT INTO swipes(user_id, profile_id, swipe_type, created_at) VALUES (?, ?, ?, ?);`
	exec, err := s.db.Exec(ctx, query, data.UserID, data.ProfileID, 1, data.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *swipeRepository) SwipeLeft(ctx context.Context, data entity.Swipe) (int64, error) {
	query := `INSERT INTO swipes(user_id, profile_id, swipe_type, created_at) VALUES (?, ?, ?, ?);`
	exec, err := s.db.Exec(ctx, query, data.UserID, data.ProfileID, 0, data.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewSwipeRepository(db databasex.Database) repository.Swipe {
	return &swipeRepository{db: db}
}
