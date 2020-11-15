package currency

import (
	"errors"
	"github.com/maratkanov-a/bank/pkg/accounts"
)

const (
	USD = "USD"
	EUR = "EUR"
	RU  = "RU"
)

var ErrorUnknownCurrency = errors.New("unknown currency")

func ConvertCurrencyToProto(c string) (accounts.CurrencyType, error) {
	switch c {
	case USD:
		return accounts.CurrencyType_USD, nil
	case EUR:
		return accounts.CurrencyType_EUR, nil
	case RU:
		return accounts.CurrencyType_RU, nil
	}

	return -1, ErrorUnknownCurrency
}

func ConvertCurrencyToRepository(c accounts.CurrencyType) (string, error) {
	switch c {
	case accounts.CurrencyType_USD:
		return USD, nil
	case accounts.CurrencyType_EUR:
		return EUR, nil
	case accounts.CurrencyType_RU:
		return RU, nil
	}

	return "", ErrorUnknownCurrency
}
