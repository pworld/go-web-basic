package services

import (
	"context"
	"errors"
	"reflect"
)

func Populate(target interface{}) error {
	return PopulateForContext(context.Background(), target)
}
func PopulateForContext(c context.Context, target interface{}) (err error) {
	return PopulateForContextWithAddVars(c, target,
		make(map[reflect.Type]reflect.Value))
}

func PopulateForContextWithAddVars(c context.Context, target interface{},
	addVars map[reflect.Type]reflect.Value) (err error) {
	targetValue := reflect.ValueOf(target)

	// Any fields type is not an interface or there is no service are skipped.
	if targetValue.Kind() == reflect.Ptr && targetValue.Elem().Kind() == reflect.Struct {
		targetValue = targetValue.Elem()

		for i := 0; i < targetValue.Type().NumField(); i++ {
			fieldVal := targetValue.Field(i)
			if fieldVal.CanSet() {
				if extra, ok := addVars[fieldVal.Type()]; ok {
					fieldVal.Set(extra)
				} else {
					err := resolveServiceValue(c, fieldVal.Addr())
					if err != nil {
						return err
					}
				}
			}
		}

	} else {
		err = errors.New("type cannot be used as target")
	}
	return
}
