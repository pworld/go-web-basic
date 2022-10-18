package main

import (
	"go-web-platform/log"
)

func writeMessage(logger log.DataLogger) {
	logger.Info("Hello, Platform")
}

func main() {
	var logger log.DataLogger = log.NewDefaultLog(log.Info)
	writeMessage(logger)
}
