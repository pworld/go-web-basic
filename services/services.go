package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifetime
}

var services = make(map[reflect.Type]BindingMap)

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

// Register Services by creating BindingMap and adding to the 'services' variable
func registerService(lt lifetime, factoryFunc interface{}) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifetime:    lt,
		}
	} else {
		err = fmt.Errorf("type cannot be used as service: %v", factoryFuncType)
	}
	return
}

func resolveServiceValue(c context.Context, val reflect.Value) (err error) {
	serviceType := val.Elem().Type()
	if serviceType == contextReferenceType {
		val.Elem().Set(reflect.ValueOf(c))
	} else if binding, found := services[serviceType]; found {
		if binding.lifetime == Scoped {
			err := resolveScopedServiceValue(c, val, binding)
			if err != nil {
				return err
			}
		} else {
			val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
		}
	} else {
		err = fmt.Errorf("cannot find service %v", serviceType)
	}
	return
}

func resolveScopedServiceValue(c context.Context, val reflect.Value,
	binding BindingMap) (err error) {
	sm, ok := c.Value(ServiceKey).(servicesMap)
	if ok {
		serviceVal, ok := sm[val.Type()]
		if !ok {
			serviceVal = invokeFunction(c, binding.factoryFunc)[0]
			sm[val.Type()] = serviceVal
		}
		val.Elem().Set(serviceVal)
	} else {
		val.Elem().Set(invokeFunction(c, binding.factoryFunc)[0])
	}
	return
}

func invokeFunction(c context.Context, f reflect.Value, otherArgs ...interface{}) []reflect.Value {
	return f.Call(resolveFunctionArgs(c, f, otherArgs...))
}

func resolveFunctionArgs(c context.Context, f reflect.Value, otherArgs ...interface{}) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := resolveServiceValue(c, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}
