package repos

import (
	"testing"
	"time"

	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	zeroTime = time.Time{}.In(time.UTC)
)

func TestAccountsList(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repo; expect empty ", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		actual, err := accountsRepo.List(newCtx())
		require.NoError(t, err)
		require.Empty(t, actual)
	})

	t.Run("one item; expect one ", func(t *testing.T) {
		expected := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, expected)
		defer testDB.Teardown()

		actual, err := accountsRepo.List(newCtx())
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		for _, a := range actual {
			a.CreatedAt = zeroTime
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("3 item; expect 3 ", func(t *testing.T) {
		expected := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, expected)
		defer testDB.Teardown()

		actual, err := accountsRepo.List(newCtx())
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		for _, a := range actual {
			a.CreatedAt = zeroTime
		}

		assert.Equal(t, expected, actual)
	})
}

func TestListByAvailability(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repo; expect empty ", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		actual, err := accountsRepo.List(newCtx())
		require.NoError(t, err)
		require.Empty(t, actual)
	})

	t.Run("3 unavailable; expect 3 ", func(t *testing.T) {
		expected := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: false, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false, CreatedAt: zeroTime},
		}
		testDB.Setup(t, expected)
		defer testDB.Teardown()

		actual, err := accountsRepo.List(newCtx())
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		for _, a := range actual {
			a.CreatedAt = zeroTime
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("3 available item; expect 3 ", func(t *testing.T) {
		expected := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, expected)
		defer testDB.Teardown()

		actual, err := accountsRepo.ListByAvailability(newCtx(), true)
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		for _, a := range actual {
			a.CreatedAt = zeroTime
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("1 available, 2 unavailable; expect 3 ", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false, CreatedAt: zeroTime},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		actual, err := accountsRepo.ListByAvailability(newCtx(), true)
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		for _, a := range actual {
			a.CreatedAt = zeroTime
		}

		expected := []*repository.Account{
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
		}

		assert.Equal(t, expected, actual)
	})
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

	t.Run("unknown id, expected error", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false, CreatedAt: zeroTime},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		_, err := accountsRepo.GetByID(newCtx(), 1111)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false, CreatedAt: zeroTime},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		actual, err := accountsRepo.GetByID(newCtx(), 11)
		require.NoError(t, err)

		expected := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime}
		actual.CreatedAt = zeroTime
		assert.Equal(t, expected, actual)
	})
}

func TestAccountsCreate(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("create incorrect currency; expected error", func(t *testing.T) {
		actual := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "UNKNOWN", IsAvailable: true, CreatedAt: zeroTime}

		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsRepo.Create(newCtx(), actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input value for enum currency_type_enum")
	})

	t.Run(" expect ok", func(t *testing.T) {
		expected := &repository.Account{Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime}

		testDB.Setup(t)
		defer testDB.Teardown()

		id, err := accountsRepo.Create(newCtx(), expected)
		require.NoError(t, err)
		require.NotZero(t, id)

		actual, err := accountsRepo.GetByID(newCtx(), id)
		require.NoError(t, err)

		expected.ID = id
		actual.CreatedAt = zeroTime
		assert.Equal(t, expected, actual)
	})
}

func TestAccountsUpdate(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		err := accountsRepo.Update(newCtx(), &repository.Account{ID: 1})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("unknown id, expected error", func(t *testing.T) {
		fixture := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime}
		testDB.Setup(t, fixture)
		defer testDB.Teardown()

		err := accountsRepo.Update(newCtx(), &repository.Account{ID: 1111})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixtures := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		expected := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime}
		err := accountsRepo.Update(newCtx(), expected)
		require.NoError(t, err)

		actual, err := accountsRepo.GetByID(newCtx(), expected.ID)
		require.NoError(t, err)

		actual.CreatedAt = zeroTime
		assert.Equal(t, expected, actual)
	})
}

func TestAccountsDelete(t *testing.T) {
	var accountsRepo = postgresql.NewAccounts(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		err := accountsRepo.Delete(newCtx(), 1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixture := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false, CreatedAt: zeroTime}
		testDB.Setup(t, fixture)
		defer testDB.Teardown()

		err := accountsRepo.Delete(newCtx(), 11)
		require.NoError(t, err)

		actual, err := accountsRepo.GetByID(newCtx(), 11)
		require.Error(t, err)
		require.Nil(t, actual)

		assert.Contains(t, err.Error(), "not found")
	})
}
