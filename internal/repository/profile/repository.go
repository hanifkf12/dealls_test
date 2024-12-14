package profile

import (
	"context"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/pkg/databasex"
)

type profileRepository struct {
	db databasex.Database
}

func (p *profileRepository) FindAll(ctx context.Context, userID int, gender string) ([]entity.Profile, error) {
	query := `SELECT * FROM profiles WHERE profiles.id NOT IN (SELECT swipes.profile_id FROM swipes WHERE swipes.user_id=? AND DATE(swipes.created_at) = CURDATE()) AND profiles.user_id != ? AND profiles.gender != ?;`
	var result = make([]entity.Profile, 0)
	err := p.db.Select(ctx, &result, query, userID, userID, gender)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *profileRepository) FindByUsersID(ctx context.Context, userID int) (*entity.Profile, error) {
	query := `SELECT * FROM profiles WHERE profiles.user_id = ?;`
	var result entity.Profile
	err := p.db.Get(ctx, &result, query, userID)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

func (p *profileRepository) Create(ctx context.Context, profile entity.Profile) (int64, error) {
	query := `INSERT INTO profiles(user_id, name, avatar, age, gender, bio, location, created_at, updated_at) VALUES (?, ?, ?, ?,?,?,?,?,?);`
	result, err := p.db.Exec(ctx, query, profile.UserID, profile.Name, profile.Avatar, profile.Age, profile.Gender, profile.Bio, profile.Location, profile.CreatedAt, profile.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *profileRepository) Update(ctx context.Context, profile entity.Profile) error {
	query := `UPDATE profiles SET name = ?, avatar = ?, age = ?, gender = ?, bio = ?, location = ?, updated_at = ? WHERE id = ?;`

	_, err := p.db.Exec(ctx, query, profile.Name, profile.Avatar, profile.Age, profile.Gender, profile.Bio, profile.Location, profile.UpdatedAt, profile.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *profileRepository) Delete(ctx context.Context, id int) error {
	query := `UPDATE profiles SET deleted_at = ? WHERE id = ?;`
	_, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func NewProfileRepository(db databasex.Database) repository.Profile {
	return &profileRepository{
		db: db,
	}
}
