package repository

import "errors"

var (
	ErrObjectNotFound       = errors.New("object not found")
	ErrNoBalance            = errors.New("insufficient balance")
	ErrIncompatibleCurrency = errors.New("incompatible currency")
)
