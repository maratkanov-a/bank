package repository

import (
	"context"
)

type AccountRepository interface {
	List(ctx context.Context) ([]*Account, error)
	ListByAvailability(ctx context.Context, isAvailable bool) ([]*Account, error)

	GetByID(ctx context.Context, id int64) (*Account, error)

	Create(ctx context.Context, a *Account) (int64, error)
	Update(ctx context.Context, a *Account) error

	Delete(ctx context.Context, id int64) error
}

type PaymentRepository interface {
	List(ctx context.Context) ([]*Payment, error)
	GetByIDMock(ctx context.Context, id int64) (*Payment, error)
	Create(ctx context.Context, from, to int64, amount int64) (int64, error)
}
