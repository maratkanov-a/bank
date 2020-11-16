package repos

import (
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"sync"
	"testing"

	"github.com/maratkanov-a/bank/internal/pkg/repository/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentsList(t *testing.T) {
	var paymentsRepo = postgresql.NewPayments(testDB.DB)

	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		actual, err := paymentsRepo.List(newCtx())
		require.NoError(t, err)
		require.Empty(t, actual)
	})

	t.Run("2 payments; expect 2", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		payments := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, accounts, payments)
		defer testDB.Teardown()

		actual, err := paymentsRepo.List(newCtx())
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		expected := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		for _, a := range actual {
			a.CreatedAt = zeroTime
		}
		assert.Equal(t, expected, actual)
	})
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

	t.Run("get incorrect; expected error", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		payments := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, accounts, payments)
		defer testDB.Teardown()

		_, err := paymentsRepo.GetByID(newCtx(), 1)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("expect ok", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		payments := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, accounts, payments)
		defer testDB.Teardown()

		actual, err := paymentsRepo.GetByID(newCtx(), 11)
		require.NoError(t, err)
		require.NotEmpty(t, actual)

		expected := &repository.Payment{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"}
		actual.CreatedAt = zeroTime
		assert.Equal(t, expected, actual)
	})
}

func TestPaymentsCreate(t *testing.T) {
	var (
		paymentsRepo = postgresql.NewPayments(testDB.DB)
		accountsRepo = postgresql.NewAccounts(testDB.DB)
	)

	t.Run("bad balance; expect error", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, accounts)
		defer testDB.Teardown()

		_, err := paymentsRepo.Create(newCtx(), 11, 12, 12)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "insufficient balance")

	})

	t.Run("different currency; expect error", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, accounts)
		defer testDB.Teardown()

		_, err := paymentsRepo.Create(newCtx(), 11, 12, 11)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "incompatible currency")
	})

	t.Run("expect ok", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, accounts)
		defer testDB.Teardown()

		id, err := paymentsRepo.Create(newCtx(), 11, 12, 11)
		require.NoError(t, err)
		require.NotZero(t, id)

		listPayments, err := paymentsRepo.List(newCtx())
		require.NoError(t, err)
		require.Len(t, listPayments, 2)

		expectedPayments := []*repository.Payment{
			{ID: id, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: id + 1, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "outgoing"},
		}
		for _, a := range listPayments {
			a.CreatedAt = zeroTime
		}
		assert.Equal(t, expectedPayments, listPayments)

		from, err := accountsRepo.GetByID(newCtx(), 11)
		require.NoError(t, err)
		expectedFrom := &repository.Account{ID: 11, Name: "11", Balance: 0, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime}
		from.CreatedAt = zeroTime
		assert.Equal(t, expectedFrom, from)

		to, err := accountsRepo.GetByID(newCtx(), 12)
		require.NoError(t, err)
		expectedTo := &repository.Account{ID: 12, Name: "22", Balance: 33, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime}
		to.CreatedAt = zeroTime
		assert.Equal(t, expectedTo, to)
	})

	t.Run("two parallel ; expect ok", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime},
		}
		testDB.Setup(t, accounts)
		defer testDB.Teardown()

		var wg = sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()
			id, err := paymentsRepo.Create(newCtx(), 11, 12, 11)
			require.NoError(t, err)
			require.NotZero(t, id)
		}()

		go func() {
			defer wg.Done()
			id, err := paymentsRepo.Create(newCtx(), 11, 12, 11)
			require.NoError(t, err)
			require.NotZero(t, id)
		}()

		wg.Wait()

		listPayments, err := paymentsRepo.List(newCtx())
		require.NoError(t, err)
		require.Len(t, listPayments, 4)

		from, err := accountsRepo.GetByID(newCtx(), 11)
		require.NoError(t, err)
		expectedFrom := &repository.Account{ID: 11, Name: "11", Balance: 0, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime}
		from.CreatedAt = zeroTime
		assert.Equal(t, expectedFrom, from)

		to, err := accountsRepo.GetByID(newCtx(), 12)
		require.NoError(t, err)
		expectedTo := &repository.Account{ID: 12, Name: "22", Balance: 44, Currency: "RU", IsAvailable: true, CreatedAt: zeroTime}
		to.CreatedAt = zeroTime
		assert.Equal(t, expectedTo, to)
	})
}
