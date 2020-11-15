package accounts

import (
	"github.com/maratkanov-a/bank/internal/pkg/balance"
	"github.com/maratkanov-a/bank/internal/pkg/currency"
	"github.com/maratkanov-a/bank/internal/pkg/repository"
	desc "github.com/maratkanov-a/bank/pkg/accounts"
	"github.com/sirupsen/logrus"
	"github.com/utrack/clay/v2/transport"
)

type Implementation struct {
	ar repository.AccountRepository
}

// NewAccounts create new Implementation
func NewAccounts(ar repository.AccountRepository) *Implementation {
	return &Implementation{
		ar: ar,
	}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *Implementation) GetDescription() transport.ServiceDesc {
	return desc.NewAccountsServiceDesc(i)
}

func convertToProtos(acs []*repository.Account) []*desc.Account {
	var converted = make([]*desc.Account, 0, len(acs))
	for _, ac := range acs {
		convertedAC, err := convertToProto(ac)
		if err != nil {
			logrus.Error(err)
			continue
		}

		converted = append(converted, convertedAC)
	}

	return converted
}

func convertToProto(ac *repository.Account) (*desc.Account, error) {
	c, err := currency.ConvertCurrencyToProto(ac.Currency)
	if err != nil {
		return nil, err
	}

	cb, err := balance.ConvertFromCents(ac.Balance)
	if err != nil {
		return nil, err
	}

	return &desc.Account{
		ID:          ac.ID,
		Name:        ac.Name,
		Balance:     cb,
		Currency:    c,
		IsAvailable: ac.IsAvailable,
	}, nil
}
