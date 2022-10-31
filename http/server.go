package http

import (
	"fmt"
	"go-web-platform/config"
	"go-web-platform/logs"
	"go-web-platform/middleware"
	"net/http"
	"sync"
)

type pipelineAdaptor struct {
	middleware.ReqPipeline
}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	err := p.ProcessRequest(request, writer)
	if err != nil {
		return
	}
}

// Http or Https Process
func Serve(pl middleware.ReqPipeline, cfg config.Config, logger logs.Logger) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	adaptor := pipelineAdaptor{ReqPipeline: pl}
	enableHttp := cfg.GetBoolDefaultValue("http:enableHttp", true)

	if enableHttp {
		httpPort := cfg.GetIntDefaultValue("http:port", 5000)
		logger.Debugf("Starting HTTP server on port %v", httpPort)
		wg.Add(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
			if err != nil {
				panic(err)
			}
		}()
	}

	enableHttps := cfg.GetBoolDefaultValue("http:enableHttps", false)

	if enableHttps {
		httpsPort := cfg.GetIntDefaultValue("http:httpsPort", 5500)
		certFile, cfok := cfg.GetString("http:httpsCert")
		keyFile, kfok := cfg.GetString("http:httpsKey")
		if cfok && kfok {
			logger.Debugf("Starting HTTPS server on port %v", httpsPort)
			wg.Add(1)
			go func() {
				err := http.ListenAndServeTLS(fmt.Sprintf(":%v", httpsPort),
					certFile, keyFile, adaptor)
				if err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS certificate settings not found")
		}
	}
	return &wg
}
