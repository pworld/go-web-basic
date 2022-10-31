package middleware

import (
	"net/http"
)

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (m *ComponentContext) Error(err error) {
	m.error = err
}
func (m *ComponentContext) GetError() error {
	return m.error
}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(context *ComponentContext, next func(*ComponentContext))
}

// Func for components to indicate they require services using reflect
type ServicesMiddlwareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}
