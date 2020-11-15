package repos

import (
	"testing"

	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountsList(t *testing.T) {
	//	TODO
}

func TestListByAvailability(t *testing.T) {
	//	TODO
}

func TestAccountsGetByID(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsRepo.GetByID(newCtx(), 1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}

func TestAccountsCreate(t *testing.T) {
	//	TODO
}

func TestAccountsUpdate(t *testing.T) {
	//	TODO
}

func TestAccountsDelete(t *testing.T) {
	//	TODO
}
