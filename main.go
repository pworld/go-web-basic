package main

import (
	"go-web-platform/config"
	"go-web-platform/log"
)

func writeMessage(logger log.DataLogger, cfg config.Config) {
	nested, ok := cfg.GetNested("main")
	if ok {
		message, ok := nested.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	var cfg config.Config
	var err error
	cfg, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}

	var logger = log.NewDefaultLog(cfg)
	writeMessage(logger, cfg)
}
