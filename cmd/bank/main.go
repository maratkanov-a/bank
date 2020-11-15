package main

import (
	"github.com/golang/glog"
	"github.com/maratkanov-a/bank/internal/app/accounts"
	"github.com/maratkanov-a/bank/internal/app/payments"
	"github.com/maratkanov-a/bank/internal/pkg/config"
	"github.com/maratkanov-a/bank/internal/pkg/server"
	"github.com/sirupsen/logrus"
	_ "github.com/utrack/clay/doc/example/static/statik"
	"github.com/utrack/clay/v2/transport"
)

func run() error {
	cfg, err := config.GetEnv()
	if err != nil {
		logrus.Fatalf("can't get environments: %v", err)
	}

	//db, err := database.NewDB(database.Options{
	//	User:            cfg.DatabaseUser,
	//	Password:        cfg.DatabasePassword,
	//	DBName:          cfg.DatabaseDBName,
	//	Host:            cfg.DatabaseHost,
	//	Port:            cfg.DatabasePort,
	//	MaxIdleConns:    cfg.DatabaseMaxIdleConns,
	//	MaxOpenConns:    cfg.DatabaseMaxOpenConns,
	//	ConnMaxLifetime: cfg.DatabaseConnMaxLifetime,
	//})
	//if err != nil {
	//	log.Fatalf("can't connect to database: %v", err)
	//}

	//accountsRepo := postgresql.NewAccounts(db)
	//paymentsRepo := postgresql.NewPayments(db)

	accountsClient := accounts.NewAccounts()
	paymentsClient := payments.NewPayments()

	compound := transport.NewCompoundServiceDesc(
		accountsClient.GetDescription(),
		paymentsClient.GetDescription(),
	)

	return server.RunAll(compound, cfg)
}

func main() {
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
