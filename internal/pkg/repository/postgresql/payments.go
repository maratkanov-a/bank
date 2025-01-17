package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/maratkanov-a/bank/internal/pkg/db"
	"github.com/maratkanov-a/bank/internal/pkg/direction"
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

func (r *PaymentsRepo) List(ctx context.Context) ([]*repository.Payment, error) {
	result := make([]*repository.Payment, 0)
	if err := r.db.SelectContext(ctx, &result, "SELECT * FROM payment"); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PaymentsRepo) GetByID(ctx context.Context, id int64) (*repository.Payment, error) {
	var payment repository.Payment
	err := r.db.QueryRowxContext(ctx, `
		SELECT
			* 
		FROM 
			payment
		WHERE 
			id = $1
	`, id).StructScan(&payment)

	if err == sql.ErrNoRows {
		return nil, repository.ErrObjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func createPayment(ctx context.Context, db database.Tx, p *repository.Payment) (int64, error) {
	var id int64

	err := db.QueryRowContext(ctx, `
		INSERT INTO payment(
			amount,
			account_from,
			account_to,
			direction,
			created_at) 
		VALUES($1,$2,$3,$4,$5) 
		RETURNING id`,
		p.Amount,
		p.AccountFrom,
		p.AccountTo,
		p.Direction,
		time.Now(),
	).Scan(&id)

	return id, err

}

func createPayments(ctx context.Context, db database.Tx, toID, fromID, amount int64) (int64, error) {
	// create to
	id, err := createPayment(ctx, db, &repository.Payment{
		Amount:      amount,
		AccountFrom: fromID,
		AccountTo:   toID,
		Direction:   direction.Incoming,
	})
	if err != nil {
		return 0, err
	}

	// create from
	_, err = createPayment(ctx, db, &repository.Payment{
		Amount:      amount,
		AccountFrom: toID,
		AccountTo:   fromID,
		Direction:   direction.Outgoing,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func updateBalances(ctx context.Context, db database.Tx, at, af *repository.Account, amount int64) error {
	ok, err := updateLockedBalance(ctx, db, at.ID, at.Balance+amount)
	if err != nil {
		return err
	}

	if !ok {
		return repository.ErrObjectNotFound
	}

	ok, err = updateLockedBalance(ctx, db, af.ID, af.Balance-amount)
	if err != nil {
		return err
	}

	if !ok {
		return repository.ErrObjectNotFound
	}

	return nil
}

func (r *PaymentsRepo) Create(ctx context.Context, from, to int64, amount int64) (int64, error) {
	var id int64
	err := r.db.WithTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}, func(ctx context.Context, tx database.Tx) error {
		accountFrom, err := getAccountByIDLocked(ctx, tx, from)
		if err != nil {
			return err
		}

		if accountFrom.Balance-amount < 0 {
			return repository.ErrNoBalance
		}

		accountTo, err := getAccountByIDLocked(ctx, tx, to)
		if err != nil {
			return err
		}

		if accountTo.Currency != accountFrom.Currency {
			return repository.ErrIncompatibleCurrency
		}

		id, err = createPayments(ctx, tx, accountTo.ID, accountFrom.ID, amount)
		if err != nil {
			return err
		}

		return updateBalances(ctx, tx, accountTo, accountFrom, amount)
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
