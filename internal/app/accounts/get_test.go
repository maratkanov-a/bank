// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: accounts.proto

package accounts

import (
	"context"
	"errors"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImplementation_Get(t *testing.T) {
	var (
		ctx             = context.Background()
		ID        int64 = 1192
		someError       = errors.New("some error")
		validReq        = &accounts.GetRequest{ID: ID}

		expectAccountRepo = &repository.Account{
			ID:       ID,
			Name:     "someName",
			Balance:  10011,
			Currency: "RU",
		}

		invalidEAccountRepo = &repository.Account{
			ID:       ID,
			Name:     "someName",
			Balance:  -1,
			Currency: "RU",
		}
	)

	t.Run("validate", func(t *testing.T) {
		for _, tc := range []struct {
			name string

			req *accounts.GetRequest

			errorMessage string
		}{
			{
				name:         "invalid name; expect error",
				req:          &accounts.GetRequest{},
				errorMessage: "invalid GetRequest.ID",
			},
		} {
			t.Run(tc.name, func(t *testing.T) {
				i := Implementation{}

				resp, err := i.Get(ctx, tc.req)
				require.Error(t, err)
				require.Nil(t, resp)

				assert.Contains(t, err.Error(), tc.errorMessage)
			})
		}
	})

	t.Run("repo err; expect err", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.GetByIDMock.Expect(ctx, ID).Return(nil, someError)

		resp, err := i.Get(ctx, validReq)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Equal(t, someError, err)
	})

	t.Run("convert err; expect error", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.GetByIDMock.Return(invalidEAccountRepo, nil)

		resp, err := i.Get(ctx, validReq)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "can't convert 0.-1 to decimal")
	})

	t.Run("expect ok", func(t *testing.T) {
		i := newTestImplementation(t)
		i.arMock.GetByIDMock.Return(expectAccountRepo, nil)

		resp, err := i.Get(ctx, validReq)
		require.NoError(t, err)
		require.NotNil(t, resp)

		converted, err := convertToProto(expectAccountRepo)
		require.NoError(t, err)

		assert.Equal(t, converted, resp.Account)
	})
}
