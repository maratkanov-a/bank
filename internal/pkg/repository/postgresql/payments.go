package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maratkanov-a/bank/internal/pkg/db"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
)

// PaymentsRepo is implementation for PostgreSQL
type PaymentsRepo struct {
	db database.DB
}

// NewPayments creates new repository for provided DB
func NewPayments(db database.DB) *PaymentsRepo {
	return &PaymentsRepo{
		db: db,
	}
}

func (r *PaymentsRepo) List(ctx context.Context, ids []int64) ([]*repository.Payment, error) {
	result := make([]*repository.Payment, 0)

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
