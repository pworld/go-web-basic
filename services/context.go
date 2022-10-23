package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type servicesMap map[reflect.Type]reflect.Value

// NewServiceContext Store RESOLVED services that have been resolved to map.
func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(servicesMap))
	}
	return c
}
