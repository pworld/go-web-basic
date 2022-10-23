package main

import (
	"go-web-platform/config"
	"go-web-platform/log"
	"go-web-platform/services"
)

func writeMessage(logger log.DataLogger, cfg config.Config) {
	section, ok := cfg.GetNested("main")
	if ok {
		message, ok := section.GetString("message")
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
	services.RegisterDefaultServices()
	_, err := services.CallFunc(writeMessage)
	if err != nil {
		return
	}
	val := struct {
		message string
		log.DataLogger
	}{
		message: "Hello from the struct",
	}
	errPop := services.Populate(&val)
	if errPop != nil {
		return
	}
	val.DataLogger.Debug(val.message)
}
