package direction

import (
	"errors"
	"github.com/maratkanov-a/bank/pkg/payments"
)

const (
	Incoming = "incoming"
	Outgoing = "outgoing"
)

var ErrorUnknownDirection = errors.New("unknown direction")

func ConvertCurrencyToProto(c string) (payments.DirectionType, error) {
	switch c {
	case Incoming:
		return payments.DirectionType_incoming, nil
	case Outgoing:
		return payments.DirectionType_outgoing, nil
	}

	return -1, ErrorUnknownDirection
}

func ConvertToRepository(c payments.DirectionType) (string, error) {
	switch c {
	case payments.DirectionType_incoming:
		return Incoming, nil
	case payments.DirectionType_outgoing:
		return Outgoing, nil
	}

	return "", ErrorUnknownDirection
}
