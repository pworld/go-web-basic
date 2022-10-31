package view

import (
	"go-web-platform/http"
	"go-web-platform/middleware"
	"go-web-platform/middleware/basic"
	"go-web-platform/services"
	"sync"
)

// Function to creates a pipeline with the middleware components created previously
func createPipeline() middleware.ReqPipeline {
	return middleware.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&SimpleMessageComponent{},
	)
}

// Func to configure and start the HTTP server
func Start() {
	results, err := services.CallFunc(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
