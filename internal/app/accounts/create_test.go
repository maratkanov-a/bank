// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: accounts.proto

package accounts

import (
	"context"
	"testing"

	desc "github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/stretchr/testify/require"
)

func TestImplementation_Create(t *testing.T) {
	api := NewAccounts()
	_, err := api.Create(context.Background(), &desc.CreateRequest{})

	require.NotNil(t, err)
	require.Equal(t, "Create not implemented", err.Error())
}
