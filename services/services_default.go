package services

import (
	"go-web-platform/config"
	"go-web-platform/log"
)

// Register Services here using Dependency Injection Method
func RegisterDefaultServices() {

	// Creates Config Services
	err := RegisterSingleton(func() (c config.Config) {
		_, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})

	// Creates Log Services
	err = RegisterSingleton(func(appConfig config.Config) log.DataLogger {
		return log.NewDefaultLog(appConfig)
	})
	if err != nil {
		panic(err)
	}
}
