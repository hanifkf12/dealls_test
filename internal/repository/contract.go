package repository

import (
	"context"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
)

type HomeRepository interface {
	GetAdmin(ctx context.Context, data any) ([]entity.Admin, error)
}

type Profile interface {
	FindAll(ctx context.Context, userID int, gender string) ([]entity.Profile, error)
	FindByUsersID(ctx context.Context, userID int) (*entity.Profile, error)
	Create(ctx context.Context, profile entity.Profile) (int64, error)
	Update(ctx context.Context, profile entity.Profile) error
	Delete(ctx context.Context, id int) error
}

type User interface {
	Create(ctx context.Context, user entity.User) (int64, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type Swipe interface {
	SwipeRight(ctx context.Context, data entity.Swipe) (int64, error)
	SwipeLeft(ctx context.Context, data entity.Swipe) (int64, error)
	IsLimit(ctx context.Context, userID int) (bool, error)
}

type Transaction interface {
	CreateTransaction(ctx context.Context, data entity.Transaction) (int64, error)
	CheckPremiumStatus(ctx context.Context, userID int) (bool, error)
}
