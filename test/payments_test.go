package integration_test

import (
	"testing"

	"github.com/maratkanov-a/bank/pkg/payments"
	"github.com/stretchr/testify/assert"
)

func TestPaymentsList(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsClient.List(newCtx(), &payments.ListRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestPaymentsGet(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsClient.Get(newCtx(), &payments.GetRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestPaymentsCreate(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}
