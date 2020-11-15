package integration_test

import (
	"testing"

	"github.com/maratkanov-a/bank/pkg/payments"
	"github.com/stretchr/testify/assert"
)

func TestPaymentsGet(t *testing.T) {
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsClient.Get(newCtx(), &payments.GetRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}
