package testdb

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/maratkanov-a/bank/internal/pkg/repository"
)

type namedExecer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type inserter struct {
	db namedExecer
}

func newInserter(db namedExecer) *inserter {
	return &inserter{db: db}
}

func (ins *inserter) namedExec(ctx context.Context, q string, fieldSlice []string, data interface{}) error {
	fieldsMap := rebindSlices(data)
	args := make([]interface{}, len(fieldSlice))
	for i, field := range fieldSlice {
		args[i] = fieldsMap[field]
	}

	_, err := ins.db.ExecContext(ctx, q, args...)
	return err
}

func (ins *inserter) insertObj(ctx context.Context, obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			err := ins.insertSingleObj(ctx, val.Index(i))
			if err != nil {
				return err
			}
		}
		return nil
	}
	return ins.insertSingleObj(ctx, val)
}

func valueToInterface(val reflect.Value) interface{} {
	if val.Kind() == reflect.Ptr {
		return val.Elem().Interface()
	}
	return val.Interface()
}

func (ins *inserter) insertSingleObj(ctx context.Context, val reflect.Value) error {
	obj := valueToInterface(val)
	switch v := obj.(type) {
	case repository.Account:
		return ins.InsertAccount(ctx, v)
	case repository.Payment:
		return ins.InsertPayment(ctx, v)
	default:
		return fmt.Errorf("unknown type for inserting: %T", obj)
	}
}

func (ins *inserter) InsertAccount(ctx context.Context, v repository.Account) error {
	q := `
		INSERT INTO account(id, name, balance, currency, is_available, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)`
	fields := []string{"id", "name", "balance", "currency", "is_available", "created_at"}
	return ins.namedExec(ctx, q, fields, v)
}

func (ins *inserter) InsertPayment(ctx context.Context, v repository.Payment) error {
	q := `
		INSERT INTO payment(id, amount, account_from, account_to, direction, created_at)
			VALUES ($1,$2,$3,$4,$5,$6)`
	fields := []string{"id", "amount", "account_from", "account_to", "direction", "created_at"}
	return ins.namedExec(ctx, q, fields, v)
}
