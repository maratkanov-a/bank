package balance

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

var ErrorIncorrectValue = errors.New("incorrect currency value")

func ConvertToCents(v float64) (int64, error) {
	d := decimal.NewFromFloat(v)
	a := d.Exponent()
	if a < -2 {
		return 0, ErrorIncorrectValue
	}

	return int64(v * 100), nil
}

func ConvertFromCents(v int64) (float64, error) {
	d, err := decimal.NewFromString(fmt.Sprintf("%d.%d", v/100, v%100))
	if err != nil {
		return 0, err
	}

	fv, _ := d.Float64()

	return fv, nil
}
