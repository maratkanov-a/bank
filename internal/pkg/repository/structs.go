package repository

import "time"

type (
	Account struct {
		ID          int64     `db:"id"` // TODO: uint64
		Name        string    `db:"name"`
		Balance     int64     `db:"balance"` // TODO: uint64
		Currency    string    `db:"currency"`
		IsAvailable bool      `db:"is_available"` // TODO: когда аккаунты доступны
		CreatedAt   time.Time `db:"created_at"`
	}

	Payment struct {
		ID          int64     `db:"id"`     // TODO: uint64
		Amount      int64     `db:"amount"` // TODO: uint64
		AccountFrom int64     `db:"account_from"`
		AccountTo   int64     `db:"account_to"`
		Direction   string    `db:"direction"`
		CreatedAt   time.Time `db:"created_at"`
	}
)
