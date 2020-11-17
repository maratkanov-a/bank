package balance

import (
	"errors"
	"strconv"

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
	if v < 0 {
		return 0, ErrorIncorrectValue
	}

	var (
		digitPart       = v / 100
		stringDigitPart = strconv.FormatInt(digitPart, 10)
		strAmount       = strconv.FormatInt(v, 10)
	)

	if digitPart == 0 {
		stringDigitPart = ""
	}

	if len(strAmount) < 2 {
		strAmount = "0" + strAmount
	}

	d, err := decimal.NewFromString(stringDigitPart + "." + strAmount[len(stringDigitPart):])
	if err != nil {
		return 0, err
	}

	fv, _ := d.Float64()

	return fv, nil
}
