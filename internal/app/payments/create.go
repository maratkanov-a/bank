package payments

import (
	"context"
	"errors"

	"github.com/maratkanov-a/bank/internal/pkg/balance"
	desc "github.com/maratkanov-a/bank/pkg/payments"
	"github.com/sirupsen/logrus"
)

var errIncorrectReceiver = errors.New("select other receiver")

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	if req.AccountTo == req.AccountFrom {
		return nil, errIncorrectReceiver
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
