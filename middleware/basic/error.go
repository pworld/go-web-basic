package basic

import (
	"fmt"
	"go-web-platform/logs"
	"go-web-platform/middleware"
	"go-web-platform/services"
	"net/http"
)

type ErrorComponent struct{}

func recoveryFunc(ctx *middleware.ComponentContext, logger logs.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

// Init recovers from  panic that occurs when subsequent components process the request and also handles any expected error.
func (c *ErrorComponent) Init() {}
func (c *ErrorComponent) ProcessRequest(ctx *middleware.ComponentContext,
	next func(*middleware.ComponentContext)) {
	var logger logs.Logger
	err := services.GetServiceContext(ctx.Context(), &logger)
	if err != nil {
		return
	}

	// Log the errors and give response status code error, defer to recovers from panics
	defer recoveryFunc(ctx, logger)
	next(ctx)
	if ctx.GetError() != nil {
		logger.Debugf("Error: %v", ctx.GetError())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
