package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DB is an interface for db ops
type DB interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	WithTx(ctx context.Context, txOpt *sql.TxOptions, f TxFunc) error
}

func newDB(driver driver.DriverContext, opts Options) (*sqlx.DB, error) {
	c, err := driver.OpenConnector(opts.ConnectionString())
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to database (%s)", opts)
	}

	db := sql.OpenDB(c)
	db.SetMaxIdleConns(opts.MaxIdleConns)
	db.SetMaxOpenConns(opts.MaxOpenConns)
	db.SetConnMaxLifetime(opts.ConnMaxLifetime)

	return sqlx.NewDb(db, "postgres"), nil
}

// NewDB creates db connection
func NewDB(opts Options) (*Database, error) {
	db, err := newDB(driverContext{driver: GetDefaultDriver()}, opts)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}
