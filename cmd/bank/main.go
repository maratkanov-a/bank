package main

import (
	"runtime"

	"github.com/maratkanov-a/bank/internal/app/accounts"
	"github.com/maratkanov-a/bank/internal/app/payments"
	"github.com/maratkanov-a/bank/internal/pkg/config"
	"github.com/maratkanov-a/bank/internal/pkg/db"
	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/maratkanov-a/bank/internal/pkg/server"
	"github.com/sirupsen/logrus"
	_ "github.com/utrack/clay/doc/example/static/statik"
	"github.com/utrack/clay/v2/transport"
)

func run() error {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	// getting data from
	cfg, err := config.GetEnv()
	if err != nil {
		logrus.Fatalf("can't get environments: %v", err)
	}

	// init connection to db
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
		logrus.Fatalf("can't connect to database: %v", err)
	}

	// init db wrappers
	accountsRepo := postgresql.NewAccounts(db)
	paymentsRepo := postgresql.NewPayments(db)

	// init services objects
	accountsClient := accounts.NewAccounts(accountsRepo)
	paymentsClient := payments.NewPayments(paymentsRepo)

	// combine services
	compound := transport.NewCompoundServiceDesc(
		accountsClient.GetDescription(),
		paymentsClient.GetDescription(),
	)

	return server.RunAll(compound, cfg)
}

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}
