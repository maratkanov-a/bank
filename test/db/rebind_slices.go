package testdb

import (
	"database/sql/driver"
	"reflect"

	"github.com/lib/pq"
)

var (
	driverValuerType = reflect.TypeOf((*driver.Valuer)(nil)).Elem()
)

// rebindSlices rebind fields to appropriate format
func rebindSlices(data interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	val := reflect.ValueOf(data)
	for i := 0; i < val.NumField(); i++ {
		ftype := val.Type().Field(i)
		fval := val.Field(i)

		fname := ftype.Tag.Get("db")
		if ftype.Type.Implements(driverValuerType) {
			result[fname] = fval.Interface()
			continue
		}

		if ftype.Type.Kind() != reflect.Slice {
			result[fname] = fval.Interface()
			continue
		}

		result[fname] = pq.Array(fval.Interface())
	}

	return result
}
