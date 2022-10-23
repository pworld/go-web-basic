package services

import (
	"context"
	"errors"
	"reflect"
)

// GetService Resolves a service using the background context
func GetService(target interface{}) error {
	return GetServiceContext(context.Background(), target)
}

// GetServiceContext accepts a context and a pointer to a value that can be set using reflection
func GetServiceContext(c context.Context, target interface{}) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr &&
		targetValue.Elem().CanSet() {
		err = resolveServiceValue(c, targetValue)
	} else {
		err = errors.New("type cannot be used as target")
	}
	return
}
