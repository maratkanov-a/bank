package repository

import "context"

type AccountRepository interface {
	List(ctx context.Context) ([]*Account, error)
	ListByAvailability(ctx context.Context, isAvailable bool) ([]*Account, error)

	GetByID(ctx context.Context, id int64) (*Account, error)

	Create(ctx context.Context, a *Account) (int64, error)
	Update(ctx context.Context, a *Account) error

	Delete(ctx context.Context, id int64) error
}

type PaymentRepository interface{}
