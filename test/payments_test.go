package integration_test

import (
	"sync"
	"testing"

	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/maratkanov-a/bank/pkg/payments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentsList(t *testing.T) {
	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		actual, err := paymentsClient.List(newCtx(), &payments.ListRequest{})
		require.NoError(t, err)
		require.Empty(t, actual)
	})

	t.Run("2 payments; expect 2", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true},
		}
		p := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, a, p)
		defer testDB.Teardown()

		resp, err := paymentsClient.List(newCtx(), &payments.ListRequest{})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expected := []*payments.Payment{
			{ID: 11, Amount: .11, AccountFrom: 11, AccountTo: 12, Direction: payments.DirectionType_incoming},
			{ID: 22, Amount: .11, AccountFrom: 12, AccountTo: 11, Direction: payments.DirectionType_outgoing},
		}
		assert.Equal(t, expected, resp.Payments)
	})
}

func TestPaymentsGet(t *testing.T) {
	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := paymentsClient.Get(newCtx(), &payments.GetRequest{ID: 1})
		require.Error(t, err)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("get incorrect; expected error", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true},
		}
		p := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, a, p)
		defer testDB.Teardown()

		_, err := paymentsClient.Get(newCtx(), &payments.GetRequest{ID: 1})
		require.Error(t, err)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("expect ok", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true},
		}
		p := []*repository.Payment{
			{ID: 11, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: "incoming"},
			{ID: 22, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: "outgoing"},
		}
		testDB.Setup(t, a, p)
		defer testDB.Teardown()

		resp, err := paymentsClient.Get(newCtx(), &payments.GetRequest{ID: 11})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expected := &payments.Payment{ID: 11, Amount: .11, AccountFrom: 11, AccountTo: 12, Direction: payments.DirectionType_incoming}
		assert.Equal(t, expected, resp.Payment)
	})
}

func TestPaymentsCreate(t *testing.T) {
	t.Run("bad balance; expect error", func(t *testing.T) {
		accounts := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 22, Currency: "RU", IsAvailable: true},
		}
		testDB.Setup(t, accounts)
		defer testDB.Teardown()

		resp, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{Amount: 12, AccountFrom: 11, AccountTo: 12})
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "insufficient balance")

	})

	t.Run("different currency; expect error", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 1100, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 2200, Currency: "USD", IsAvailable: true},
		}
		testDB.Setup(t, a)
		defer testDB.Teardown()

		resp, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{Amount: 11, AccountFrom: 11, AccountTo: 12})
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "incompatible currency")
	})

	t.Run("expect ok", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 1100, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 2200, Currency: "RU", IsAvailable: true},
		}
		testDB.Setup(t, a)
		defer testDB.Teardown()

		resp, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{Amount: 11, AccountFrom: 11, AccountTo: 12})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		list, err := paymentsClient.List(newCtx(), &payments.ListRequest{})
		require.NoError(t, err)
		require.Len(t, list.Payments, 2)

		expectedPayments := []*payments.Payment{
			{ID: resp.ID, Amount: 11, AccountFrom: 11, AccountTo: 12, Direction: payments.DirectionType_incoming},
			{ID: resp.ID + 1, Amount: 11, AccountFrom: 12, AccountTo: 11, Direction: payments.DirectionType_outgoing},
		}
		assert.Equal(t, expectedPayments, list.Payments)

		from, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 11})
		require.NoError(t, err)
		expectedFrom := &accounts.Account{ID: 11, Name: "11", Balance: 0, Currency: accounts.CurrencyType_RU, IsAvailable: true}
		assert.Equal(t, expectedFrom, from.Account)

		to, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 12})
		require.NoError(t, err)
		expectedTo := &accounts.Account{ID: 12, Name: "22", Balance: 33, Currency: accounts.CurrencyType_RU, IsAvailable: true}
		assert.Equal(t, expectedTo, to.Account)
	})

	t.Run("two parallel ; expect ok", func(t *testing.T) {
		a := []*repository.Account{
			{ID: 11, Name: "11", Balance: 2200, Currency: "RU", IsAvailable: true},
			{ID: 12, Name: "22", Balance: 2200, Currency: "RU", IsAvailable: true},
		}
		testDB.Setup(t, a)
		defer testDB.Teardown()

		var wg = sync.WaitGroup{}
		wg.Add(2)

		go func() {
			defer wg.Done()
			resp, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{Amount: 11, AccountFrom: 11, AccountTo: 12})
			require.NoError(t, err)
			require.NotEmpty(t, resp)
		}()

		go func() {
			defer wg.Done()
			resp, err := paymentsClient.Create(newCtx(), &payments.CreateRequest{Amount: 11, AccountFrom: 11, AccountTo: 12})
			require.NoError(t, err)
			require.NotEmpty(t, resp)
		}()

		wg.Wait()

		from, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 11})
		require.NoError(t, err)
		expectedFrom := &accounts.Account{ID: 11, Name: "11", Balance: 0, Currency: accounts.CurrencyType_RU, IsAvailable: true}
		assert.Equal(t, expectedFrom, from.Account)

		to, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 12})
		require.NoError(t, err)
		expectedTo := &accounts.Account{ID: 12, Name: "22", Balance: 44, Currency: accounts.CurrencyType_RU, IsAvailable: true}
		assert.Equal(t, expectedTo, to.Account)
	})
}
