package repository

import (
	"context"
)

type (
	// AccountRepository works with accounts data layer
	AccountRepository interface {
		List(ctx context.Context) ([]*Account, error)
		ListByAvailability(ctx context.Context, isAvailable bool) ([]*Account, error)
		GetByID(ctx context.Context, id int64) (*Account, error)
		Create(ctx context.Context, ac *Account) (int64, error)
		Update(ctx context.Context, ac *Account) error
		Delete(ctx context.Context, id int64) error
	}

	// PaymentRepository works with payments data layer
	PaymentRepository interface {
		List(ctx context.Context) ([]*Payment, error)
		GetByID(ctx context.Context, id int64) (*Payment, error)
		Create(ctx context.Context, from, to int64, amount int64) (int64, error)
	}
)
