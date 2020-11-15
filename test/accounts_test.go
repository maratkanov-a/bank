package integration_test

import (
	"testing"

	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/assert"
)

func TestAccountsGet(t *testing.T) {
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Get(newCtx(), &accounts.GetRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}
