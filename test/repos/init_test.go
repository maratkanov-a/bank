package repos

import (
	"context"

	testdb "github.com/maratkanov-a/bank/test/db"
)

var (
	testDB *testdb.DB
)

func init() {
	testDB = testdb.NewFromEnv()
}

func newCtx() context.Context {
	return context.Background()
}
