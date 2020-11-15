package testdb

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	database "github.com/maratkanov-a/bank/internal/pkg/db"
	"github.com/maratkanov-a/bank/test/config"
	"github.com/sirupsen/logrus"
)

// DB is concurrency-safe DB client for testing purposes
type DB struct {
	sync.Mutex
	database.DB
}

// NewFromEnv return *DB form environment variables
func NewFromEnv() *DB {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDB(database.Options{
		User:            cfg.DatabaseUser,
		Password:        cfg.DatabasePassword,
		DBName:          cfg.DatabaseDBName,
		Host:            cfg.DatabaseHost,
		Port:            cfg.DatabasePort,
		MaxIdleConns:    cfg.DatabaseMaxIdleConns,
		MaxOpenConns:    cfg.DatabaseMaxOpenConns,
		ConnMaxLifetime: cfg.DatabaseConnMaxLifetime,
	})
	if err != nil {
		panic(err)
	}

	return &DB{DB: db}
}

// Setup inserts objects into database and acquire lock, don't forget to call Teardown()
// after tests
func (db *DB) Setup(t *testing.T, objects ...interface{}) {
	ctx := context.Background()

	db.Lock()
	db.truncate(ctx)

	err := db.WithTx(ctx, nil, func(ctx context.Context, tx database.Tx) error {
		var (
			ins = newInserter(tx)
			err error
		)

		for _, obj := range objects {
			err = ins.insertObj(ctx, obj)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		db.Teardown()
		t.Fatal(err)
	}
}

func (db *DB) Insert(objects ...interface{}) error {
	ctx := context.Background()
	ins := newInserter(db.DB)
	for _, obj := range objects {
		err := ins.insertObj(ctx, obj)
		if err != nil {
			return err
		}
	}

	return nil
}

// Teardown truncate tables and release the lock
func (db *DB) Teardown() {
	defer db.Unlock()

	db.truncate(context.Background())
}

func (db *DB) truncate(ctx context.Context) {
	var tables []string
	err := db.SelectContext(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name!='goose_db_version'")
	if err != nil {
		panic(err)
	}

	if len(tables) == 0 {
		panic("no tables found, run migrations first")
	}

	q := fmt.Sprintf("TRUNCATE %s", strings.Join(tables, ","))
	_, err = db.ExecContext(ctx, q)
	if err != nil {
		logrus.Error(ctx, err)
	}
}
