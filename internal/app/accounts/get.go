// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: accounts.proto

package accounts

import (
	"context"
	"github.com/sirupsen/logrus"

	desc "github.com/maratkanov-a/bank/pkg/accounts"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if err := req.Validate(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	account, err := i.ar.GetByID(ctx, req.ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	ca, err := convertToProto(account)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &desc.GetResponse{Account: ca}, nil
}
