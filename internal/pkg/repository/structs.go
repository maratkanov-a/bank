package repository

import "time"

type (
	Account struct {
		ID          int64     `db:"id"`
		Name        string    `db:"name"`
		Balance     int64     `db:"balance"`
		Currency    string    `db:"currency"`
		IsAvailable bool      `db:"is_available"`
		CreatedAt   time.Time `db:"created_at"`
	}

	Payment struct {
		ID          int64     `db:"id"`
		Amount      int64     `db:"amount"`
		AccountFrom int64     `db:"account_from"`
		AccountTo   int64     `db:"account_to"`
		Direction   string    `db:"direction"`
		CreatedAt   time.Time `db:"created_at"`
	}
)
