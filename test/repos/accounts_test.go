package repos

import (
	"testing"

	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountsGet(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsRepo.GetByID(newCtx(), 1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}
