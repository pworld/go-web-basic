package services

import (
	"reflect"
	"sync"
)

func RegisterTransient(factoryFunc interface{}) (err error) {
	return registerService(Transient, factoryFunc)
}

func RegisterScoped(factoryFunc interface{}) (err error) {
	return registerService(Scoped, factoryFunc)
}

func RegisterSingleton(factoryFunc interface{}) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)

	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}

		// Create Wrapper around Factory Func to ensure only executed once
		wrapper := reflect.MakeFunc(factoryFuncVal.Type(),
			func([]reflect.Value) []reflect.Value {
				once.Do(func() {
					results = invokeFunction(nil, factoryFuncVal)
				})
				return results
			})
		// only one instance of the implementation struct created
		// and that it won’t be created until the first time it is needed
		err = registerService(Singleton, wrapper.Interface())
	}
	return
}
