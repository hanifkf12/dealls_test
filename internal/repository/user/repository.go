package user

import (
	"context"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/pkg/databasex"
)

type userRepository struct {
	db databasex.Database
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (int64, error) {
	query := `INSERT INTO users(email, password, created_at, updated_at) VALUES (?, ?, ?, ?);`

	result, err := u.db.Exec(ctx, query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	query := `SELECT * FROM users;`
	var results = make([]entity.User, 0)

	err := u.db.Select(ctx, &results, query)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT * FROM users WHERE email = ?;`

	var result entity.User
	err := u.db.Get(ctx, &result, query, email)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func NewUserRepository(db databasex.Database) repository.User {
	return &userRepository{db: db}
}
