package middleware

import (
	"net/http"
)

type ReqPipeline func(*ComponentContext)

var DefaultPipeline ReqPipeline = func(*ComponentContext) { /* do nothing */ }

func CreatePipeline(components ...MiddlewareComponent) ReqPipeline {
	f := DefaultPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		nextFunc := f
		f = func(context *ComponentContext) {
			if context.error == nil {
				currentComponent.ProcessRequest(context, nextFunc)
			}
		}
		currentComponent.Init()
	}
	return f
}

func (pl ReqPipeline) ProcessRequest(req *http.Request,
	resp http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: resp,
	}
	pl(&ctx)
	return ctx.error
}
