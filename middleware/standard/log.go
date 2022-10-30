package standard

import (
	"go-web-platform/logs"
	"go-web-platform/middleware"
	"go-web-platform/services"
	"net/http"
)

type LogResWriter struct {
	statusCode int
	http.ResponseWriter
}

func (w *LogResWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
func (w *LogResWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

type LoggingComponent struct{}

func (lc *LoggingComponent) Init() {}
func (lc *LoggingComponent) ProcessRequest(ctx *middleware.ComponentContext,
	next func(*middleware.ComponentContext)) {
	var logger logs.DataLogger
	err := services.GetServiceContext(ctx.Request.Context(), &logger)
	if err != nil {
		ctx.Error(err)
		return
	}
	loggingWriter := LogResWriter{0, ctx.ResponseWriter}
	ctx.ResponseWriter = &loggingWriter
	logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL)
}
