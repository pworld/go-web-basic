package basic

import (
	"go-web-platform/config"
	"go-web-platform/middleware"
	"net/http"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
	Config        config.Config
}

// Init Read Config File
func (sfc *StaticFileComponent) Init() {
	sfc.urlPrefix = sfc.Config.GetStringDefaultValue("files:urlprefix", "/files/")

	// File Config Settings
	sfc.urlPrefix = sfc.Config.GetStringDefaultValue("files:urlprefix", "/files/")
	path, ok := sfc.Config.GetString("files:path")
	if ok {
		sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix,
			http.FileServer(http.Dir(path)))
	} else {
		panic("Cannot load file configuration settings")
	}
}
func (sfc *StaticFileComponent) ProcessRequest(ctx *middleware.ComponentContext,
	next func(*middleware.ComponentContext)) {
	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
		strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
