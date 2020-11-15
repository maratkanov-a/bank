package integration_test

import (
	"context"
	"time"

	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/maratkanov-a/bank/pkg/payments"
	"github.com/maratkanov-a/bank/test/config"
	testdb "github.com/maratkanov-a/bank/test/db"
	"google.golang.org/grpc"
)

func newCtx() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_ = cancel
	return ctx
}

var (
	testDB         *testdb.DB
	paymentsClient payments.PaymentsClient
	accountsClient accounts.AccountsClient
)

func init() {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}

	appGRPCConn, err := grpc.Dial(cfg.AppAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	paymentsClient = payments.NewPaymentsClient(appGRPCConn)
	accountsClient = accounts.NewAccountsClient(appGRPCConn)

	testDB = testdb.NewFromEnv()
}
