// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: accounts.proto

package accounts

import (
	"context"
	"github.com/maratkanov-a/bank/internal/pkg/balance"
	"github.com/maratkanov-a/bank/internal/pkg/currency"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	"github.com/sirupsen/logrus"

	desc "github.com/maratkanov-a/bank/pkg/accounts"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	if err := req.Validate(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	curr, err := currency.ConvertCurrencyToRepository(req.Currency)
	if err != nil {
		return nil, err
	}

	// TODO: select for update

	bal, err := balance.ConvertToCents(req.Balance)
	if err != nil {
		return nil, err
	}

	acc := &repository.Account{
		Name:        req.Name,
		Balance:     bal,
		Currency:    curr,
		IsAvailable: req.IsAvailable,
	}

	if err := i.ar.Update(ctx, acc); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &desc.UpdateResponse{}, nil
}
