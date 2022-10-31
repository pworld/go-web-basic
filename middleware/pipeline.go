package middleware

import (
	"go-web-platform/services"
	"net/http"
	"reflect"
)

type ReqPipeline func(*ComponentContext)

var DefaultPipeline ReqPipeline = func(*ComponentContext) { /* do nothing */ }

func CreatePipeline(components ...interface{}) ReqPipeline {
	f := DefaultPipeline
	for i := len(components) - 1; i >= 0; i-- {
		currentComponent := components[i]
		err := services.Populate(currentComponent)
		if err != nil {
			return nil
		}
		nextFunc := f
		if servComp, ok := currentComponent.(ServicesMiddlwareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nextFunc)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(context *ComponentContext) {
				if context.error == nil {
					stdComp.ProcessRequest(context, nextFunc)
				}
			}
			stdComp.Init()
		} else {
			panic("Value is not a middleware component")
		}
	}
	return f
}

func createServiceDependentFunction(component interface{}, nextFunc ReqPipeline) ReqPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(context *ComponentContext) {
			if context.error == nil {
				_, err := services.CallFuncForContext(context.Request.Context(),
					method.Interface(), context, nextFunc)
				if err != nil {
					context.Error(err)
				}
			}
		}
	} else {
		panic("No ProcessRequestWithServices method defined")
	}
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
