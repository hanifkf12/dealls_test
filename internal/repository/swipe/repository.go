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

func (s *swipeRepository) IsLimit(ctx context.Context, userID int) (bool, error) {
	query := `SELECT count(id)  FROM swipes WHERE swipes.user_id=? AND DATE(swipes.created_at) = CURDATE();`
	var result int
	err := s.db.Get(ctx, &result, query, userID)
	if err != nil {
		return false, err
	}
	return result >= 10, nil
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
