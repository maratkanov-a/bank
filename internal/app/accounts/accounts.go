package accounts

import (
	desc "github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/utrack/clay/v2/transport"
)

type Implementation struct{}

// NewAccounts create new Implementation
func NewAccounts() *Implementation {
	return &Implementation{}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Implementation) GetDescription() transport.ServiceDesc {
	return desc.NewAccountsServiceDesc(i)
}
