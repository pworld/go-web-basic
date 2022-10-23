package services

import (
	"context"
	"errors"
	"reflect"
)

// CallFunc function is a convenience for use when a Context isn’t available.
func CallFunc(target interface{}, otherArgs ...interface{}) ([]interface{}, error) {
	return CallFuncForContext(context.Background(), target, otherArgs...)
}

func CallFuncForContext(c context.Context, target interface{},
	otherArgs ...interface{}) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetValue, otherArgs...)
		results = make([]interface{}, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("only functions can be invoked")
	}
	return
}
