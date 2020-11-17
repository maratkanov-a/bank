package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Tx interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
}

type TxFunc = func(context.Context, Tx) error

func (db Database) WithTx(ctx context.Context, txOpt *sql.TxOptions, f TxFunc) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	err = f(ctx, tx)

	if err == nil {
		return tx.Commit()
	}

	rbErr := tx.Rollback()
	if rbErr != nil {
		logrus.Fatal(rbErr)
	}

	return err
}
