package transaction

import (
	"context"
	"github.com/hanifkf12/hanif_skeleton/internal/entity"
	"github.com/hanifkf12/hanif_skeleton/internal/repository"
	"github.com/hanifkf12/hanif_skeleton/pkg/databasex"
)

type transactionRepository struct {
	db databasex.Database
}

func (t *transactionRepository) CreateTransaction(ctx context.Context, data entity.Transaction) (int64, error) {
	query := `INSERT INTO transactions(id, user_id, price, package_type, valid_until, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`
	exec, err := t.db.Exec(ctx, query, data.ID, data.UserID, data.Price, data.PackageType, data.ValidUntil, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *transactionRepository) CheckPremiumStatus(ctx context.Context, userID int) (bool, error) {
	query := `SELECT
       IF(valid_until >= CURDATE(), TRUE, FALSE) AS is_premium
FROM transactions
WHERE valid_until >= CURDATE() AND user_id = ?;`
	var result bool
	err := t.db.Get(ctx, &result, query, userID)
	if err != nil {
		return false, err
	}
	return result, nil
}

func NewTransactionRepository(db databasex.Database) repository.Transaction {
	return &transactionRepository{db: db}
}
