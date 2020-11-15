package postgresql

import (
	"context"
	"database/sql"
	"time"

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

func (r *AccountsRepo) List(ctx context.Context) ([]*repository.Account, error) {
	result := make([]*repository.Account, 0)
	if err := r.db.SelectContext(ctx, &result, "SELECT * FROM account"); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AccountsRepo) ListByAvailability(ctx context.Context, isAvailable bool) ([]*repository.Account, error) {
	result := make([]*repository.Account, 0)
	if err := r.db.SelectContext(ctx, &result, "SELECT * FROM account WHERE is_available = $1", isAvailable); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AccountsRepo) GetByID(ctx context.Context, id int64) (*repository.Account, error) {
	var account repository.Account
	err := r.db.QueryRowContext(ctx, `
		SELECT
			* 
		FROM 
			account 
		WHERE 
			id = $1
	`, id).Scan(&account)

	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountsRepo) Create(ctx context.Context, ac *repository.Account) (int64, error) {
	var id int64

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO account(
			name,
			balance,
			currency,
			is_available,
			created_at) 
		VALUES($1,$2,$3,$4,$5) 
		RETURNING id`,
		ac.Name,
		ac.Currency,
		ac.Balance,
		ac.IsAvailable,
		time.Now(),
	).Scan(&id)

	return id, err
}

func getAccountByIDLocked(ctx context.Context, db database.Tx, id int64) (*repository.Account, error) {
	var account repository.Account
	err := db.QueryRowContext(ctx, `
		SELECT FOR UPDATE
			* 
		FROM 
			account 
		WHERE 
			id = $1
	`, id).Scan(&account)

	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func updateAccountLocked(ctx context.Context, db database.Tx, ac *repository.Account) (bool, error) {
	result, err := db.ExecContext(ctx, `
		UPDATE
			account
		SET
			name=$2,
			balance=$3,
			currency=$4,
			is_available=$5
		WHERE
			id=$1
	`,
		ac.ID,
		ac.Name,
		ac.Currency,
		ac.Balance,
		ac.IsAvailable,
	)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (r *AccountsRepo) Update(ctx context.Context, ac *repository.Account) error {
	err := r.db.WithTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}, func(ctx context.Context, tx database.Tx) error {
		_, err := getAccountByIDLocked(ctx, tx, ac.ID)
		if err != nil {
			return err
		}

		ok, err := updateAccountLocked(ctx, tx, ac)
		if err != nil {
			return err
		}

		if !ok {
			return repository.ErrObjectNotFound
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func deleteLocked(ctx context.Context, db database.Tx, id int64) (bool, error) {
	result, err := db.ExecContext(ctx, "DELETE FROM parts WHERE id = $1", id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected > 0, err
}

func (r *AccountsRepo) Delete(ctx context.Context, id int64) error {
	err := r.db.WithTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}, func(ctx context.Context, tx database.Tx) error {
		_, err := getAccountByIDLocked(ctx, tx, id)
		if err != nil {
			return err
		}

		ok, err := deleteLocked(ctx, tx, id)
		if err != nil {
			return err
		}

		if !ok {
			return repository.ErrObjectNotFound
		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil
}
