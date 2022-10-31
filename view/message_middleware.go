package view

import (
	"errors"
	"go-web-platform/config"
	"go-web-platform/middleware"
	"go-web-platform/services"
	"io"
)

type SimpleMessageComponent struct{}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(ctx *middleware.ComponentContext,
	next func(*middleware.ComponentContext)) {
	var cfg config.Config
	err := services.GetService(&cfg)
	if err != nil {
		return
	}
	msg, ok := cfg.GetString("main:message")
	if ok {
		_, err := io.WriteString(ctx.ResponseWriter, msg)
		if err != nil {
			return
		}
	} else {
		ctx.Error(errors.New("cannot find config setting"))
	}
	next(ctx)
}
