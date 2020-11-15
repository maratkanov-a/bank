package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maratkanov-a/bank/internal/pkg/db"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
)

// AccountsRepo is implementation for PostgreSQL
type AccountsRepo struct {
	db database.DB
}

// NewAccounts creates new repository for provided DB
func NewAccounts(db database.DB) *AccountsRepo {
	return &AccountsRepo{
		db: db,
	}
}

func (r *AccountsRepo) List(ctx context.Context, ids []int64) ([]*repository.Account, error) {
	result := make([]*repository.Account, 0)

	query, args, err := sqlx.In(`
		SELECT 
			* 
		FROM 
			accounts 
		WHERE 
			id IN (?) 
		`,
		ids)
	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &result, sqlx.Rebind(sqlx.DOLLAR, query), args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
