// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: payments.proto

package payments

import (
	"context"
	"errors"
	"github.com/maratkanov-a/bank/internal/pkg/direction"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/maratkanov-a/bank/pkg/payments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestImplementation_Get(t *testing.T) {
	var (
		ctx             = context.Background()
		ID        int64 = 1192
		someError       = errors.New("some error")
		validReq        = &payments.GetRequest{ID: ID}

		accountRepo = &repository.Payment{
			ID:          ID,
			AccountFrom: 11,
			AccountTo:   12,
			Amount:      100,
			Direction:   direction.Outgoing,
		}

		expected = &payments.Payment{
			ID:          ID,
			AccountFrom: 11,
			AccountTo:   12,
			Amount:      1.00,
			Direction:   payments.DirectionType_outgoing,
		}

		invalidEAccountRepo = &repository.Payment{
			ID:          ID,
			AccountFrom: 11,
			AccountTo:   12,
			Amount:      -1,
		}
	)

	t.Run("validate", func(t *testing.T) {
		for _, tc := range []struct {
			name string

			req *payments.GetRequest

			errorMessage string
		}{
			{
				name:         "invalid name; expect error",
				req:          &payments.GetRequest{},
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
		i.prMock.GetByIDMockMock.Expect(ctx, ID).Return(nil, someError)

		resp, err := i.Get(ctx, validReq)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Equal(t, someError, err)
	})

	t.Run("convert err; expect error", func(t *testing.T) {
		i := newTestImplementation(t)
		i.prMock.GetByIDMockMock.Return(invalidEAccountRepo, nil)

		resp, err := i.Get(ctx, validReq)
		require.Error(t, err)
		require.Nil(t, resp)

		assert.Contains(t, err.Error(), "can't convert 0.-1 to decimal")
	})

	t.Run("expect ok", func(t *testing.T) {
		i := newTestImplementation(t)
		i.prMock.GetByIDMockMock.Return(accountRepo, nil)

		resp, err := i.Get(ctx, validReq)
		require.NoError(t, err)
		require.NotNil(t, resp)

		assert.Equal(t, expected, resp.Payment)
	})
}
