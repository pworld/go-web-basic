package basic

import (
	"go-web-platform/middleware"
	"go-web-platform/services"
)

type ServicesComponent struct{}

func (c *ServicesComponent) Init() {}

// The http.Request.Context method is used to get the
// basic Context created with the request, which is prepared for services and then updated using the
// WithContext method
func (c *ServicesComponent) ProcessRequest(ctx *middleware.ComponentContext,
	next func(*middleware.ComponentContext)) {
	reqContext := ctx.Request.Context()
	ctx.Request.WithContext(services.NewServiceContext(reqContext))

	//The http.Request.Context method is used to get the
	//basic Context created with the request, which is prepared for services and then updated using the
	//WithContext method
	next(ctx)
}
