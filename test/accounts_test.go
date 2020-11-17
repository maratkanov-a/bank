package integration_test

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountsList(t *testing.T) {
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		resp, err := accountsClient.List(newCtx(), &accounts.ListRequest{})
		require.NoError(t, err)
		require.Empty(t, resp)
	})

	t.Run("without availability; expect all", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: false},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: true},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		resp, err := accountsClient.List(newCtx(), &accounts.ListRequest{})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expected := []*accounts.Account{
			{ID: 11, Name: "11", Balance: .11, Currency: accounts.CurrencyType_RU, IsAvailable: true},
			{ID: 22, Name: "22", Balance: .22, Currency: accounts.CurrencyType_USD, IsAvailable: false},
			{ID: 33, Name: "33", Balance: .33, Currency: accounts.CurrencyType_EUR, IsAvailable: true},
		}
		assert.Equal(t, expected, resp.Accounts)
	})

	t.Run("list available; expect error", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: false},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: true},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		resp, err := accountsClient.List(newCtx(), &accounts.ListRequest{IsAvailable: &types.BoolValue{Value: true}})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expected := []*accounts.Account{
			{ID: 11, Name: "11", Balance: .11, Currency: accounts.CurrencyType_RU, IsAvailable: true},
			{ID: 33, Name: "33", Balance: .33, Currency: accounts.CurrencyType_EUR, IsAvailable: true},
		}
		assert.Equal(t, expected, resp.Accounts)
	})

	t.Run("list unavailable; expect error", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: true},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: false},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: true},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		resp, err := accountsClient.List(newCtx(), &accounts.ListRequest{IsAvailable: &types.BoolValue{Value: false}})
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		expected := []*accounts.Account{{ID: 22, Name: "22", Balance: .22, Currency: accounts.CurrencyType_USD, IsAvailable: false}}
		assert.Equal(t, expected, resp.Accounts)
	})
}

func TestAccountsGet(t *testing.T) {
	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 1})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("unknown id, expected error", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		_, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 1111})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixtures := []*repository.Account{
			{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false},
			{ID: 22, Name: "22", Balance: 22, Currency: "USD", IsAvailable: true},
			{ID: 33, Name: "33", Balance: 33, Currency: "EUR", IsAvailable: false},
		}
		testDB.Setup(t, fixtures)
		defer testDB.Teardown()

		resp, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 11})
		require.NoError(t, err)

		expected := &accounts.Account{ID: 11, Name: "11", Balance: .11, Currency: accounts.CurrencyType_RU, IsAvailable: false}
		assert.Equal(t, expected, resp.Account)
	})
}

func TestAccountsCreate(t *testing.T) {
	var validReq = &accounts.CreateRequest{
		Name:     "11",
		Balance:  11,
		Currency: accounts.CurrencyType_RU,
	}

	t.Run("create incorrect currency; expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Create(newCtx(), &accounts.CreateRequest{
			Name:     "11",
			Balance:  11,
			Currency: 12,
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "value must be one of the defined enum values")
	})

	t.Run(" expect ok", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		resp, err := accountsClient.Create(newCtx(), validReq)
		require.NoError(t, err)
		require.NotEmpty(t, resp)

		respGet, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: resp.ID})
		require.NoError(t, err)

		expected := &accounts.Account{ID: resp.ID, Name: "11", Balance: 11, Currency: accounts.CurrencyType_RU, IsAvailable: true}
		assert.Equal(t, expected, respGet.Account)
	})

}

func TestAccountsUpdate(t *testing.T) {
	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		resp, err := accountsClient.Update(newCtx(), &accounts.UpdateRequest{
			ID:          1,
			Name:        "11",
			Balance:     22,
			Currency:    accounts.CurrencyType_RU,
			IsAvailable: true,
		})
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("unknown id, expected error", func(t *testing.T) {
		fixture := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false}
		testDB.Setup(t, fixture)
		defer testDB.Teardown()

		resp, err := accountsClient.Update(newCtx(), &accounts.UpdateRequest{
			ID:          1111,
			Name:        "11",
			Balance:     22,
			Currency:    accounts.CurrencyType_RU,
			IsAvailable: true,
		})
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixture := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false}
		testDB.Setup(t, fixture)
		defer testDB.Teardown()

		resp, err := accountsClient.Update(newCtx(), &accounts.UpdateRequest{
			ID:          11,
			Name:        "11",
			Balance:     22,
			Currency:    accounts.CurrencyType_RU,
			IsAvailable: false,
		})
		require.NoError(t, err)
		require.Empty(t, resp)

		actual, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: fixture.ID})
		require.NoError(t, err)

		expected := &accounts.Account{ID: 11, Name: "11", Balance: 22, Currency: accounts.CurrencyType_RU, IsAvailable: false}
		assert.Equal(t, expected, actual.Account)
	})
}

func TestAccountsDelete(t *testing.T) {
	t.Run("empty repository, expected error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		resp, err := accountsClient.Delete(newCtx(), &accounts.DeleteRequest{ID: 1})
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("correct id, expected ok", func(t *testing.T) {
		fixture := &repository.Account{ID: 11, Name: "11", Balance: 11, Currency: "RU", IsAvailable: false}
		testDB.Setup(t, fixture)
		defer testDB.Teardown()

		resp, err := accountsClient.Delete(newCtx(), &accounts.DeleteRequest{ID: 11})
		require.NoError(t, err)
		require.Empty(t, resp)

		actual, err := accountsClient.Get(newCtx(), &accounts.GetRequest{ID: 11})
		require.Error(t, err)
		require.Nil(t, actual)

		assert.Contains(t, err.Error(), "not found")
	})
}
