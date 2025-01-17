// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: accounts.proto

package accounts

import (
	"context"
	"errors"
	"github.com/gogo/protobuf/types"
	"github.com/maratkanov-a/bank/internal/pkg/currency"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImplementation_List(t *testing.T) {
	var (
		ctx          = context.Background()
		someError    = errors.New("some error")
		withAvail    = &accounts.ListRequest{IsAvailable: &types.BoolValue{Value: true}}
		withoutAvail = &accounts.ListRequest{}
		repoResponse = []*repository.Account{
			{ID: 112233, Name: "112233", Balance: 112233, Currency: currency.RU},
			{ID: 445566, Name: "445566", Balance: 445566, Currency: currency.EUR},
			{ID: 778899, Name: "778899", Balance: 778899, Currency: currency.USD},
		}
		expected = []*accounts.Account{
			{ID: 112233, Name: "112233", Balance: 1122.33, Currency: accounts.CurrencyType_RU},
			{ID: 445566, Name: "445566", Balance: 4455.66, Currency: accounts.CurrencyType_EUR},
			{ID: 778899, Name: "778899", Balance: 7788.99, Currency: accounts.CurrencyType_USD},
		}
	)

	t.Run("with avail err; expect err", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.ListByAvailabilityMock.Expect(ctx, true).Return(nil, someError)

		resp, err := i.List(ctx, withAvail)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Equal(t, someError, err)
	})

	t.Run("list repo err; expect err", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.ListMock.Expect(ctx).Return(nil, someError)

		resp, err := i.List(ctx, withoutAvail)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Equal(t, someError, err)
	})

	t.Run("expect ok", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.ListMock.Return(repoResponse, nil)

		resp, err := i.List(ctx, withoutAvail)
		require.NoError(t, err)
		require.NotNil(t, resp)

		assert.Equal(t, expected, resp.Accounts)
	})
}
