package payments

import (
	"context"

	desc "github.com/maratkanov-a/bank/pkg/payments"
	"github.com/pkg/errors"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return nil, errors.New("Create not implemented")
}
