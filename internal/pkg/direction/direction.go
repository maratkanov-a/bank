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

// ConvertCurrencyToProto converts db string to proto type
func ConvertCurrencyToProto(c string) (payments.DirectionType, error) {
	switch c {
	case Incoming:
		return payments.DirectionType_incoming, nil
	case Outgoing:
		return payments.DirectionType_outgoing, nil
	}

	return -1, ErrorUnknownDirection
}
