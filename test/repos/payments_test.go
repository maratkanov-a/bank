package repos

import (
	"testing"

	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentsList(t *testing.T) {
	//	TODO
}

func TestPaymentsGetByID(t *testing.T) {
	var paymentsRepo = postgresql.NewPayments(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsRepo.GetByID(newCtx(), 1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestPaymentsCreate(t *testing.T) {
	//	TODO
}
