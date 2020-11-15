package payments

import (
	"context"
	"github.com/maratkanov-a/bank/internal/pkg/balance"
	desc "github.com/maratkanov-a/bank/pkg/payments"
	"github.com/sirupsen/logrus"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	amount, err := balance.ConvertToCents(req.Amount)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	id, err := i.pr.Create(ctx, req.AccountFrom, req.AccountTo, amount)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &desc.CreateResponse{ID: id}, nil
}
