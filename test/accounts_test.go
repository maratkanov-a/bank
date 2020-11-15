package integration_test

import (
	"testing"

	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/assert"
)

func TestAccountsList(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.List(newCtx(), &accounts.ListRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestAccountsGet(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Get(newCtx(), &accounts.GetRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestAccountsCreate(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Create(newCtx(), &accounts.CreateRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestAccountsUpdate(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Update(newCtx(), &accounts.UpdateRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}

func TestAccountsDelete(t *testing.T) {
	// TODO
	t.Run("not existing id; expect error", func(t *testing.T) {
		testDB.Setup(t)
		defer testDB.Teardown()

		_, err := accountsClient.Delete(newCtx(), &accounts.DeleteRequest{})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "object not found")
	})

}
