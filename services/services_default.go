package services

import (
	"go-web-platform/config"
	"go-web-platform/logs"
)

// Register Services here using Dependency Injection Method
func RegisterDefaultServices() {

	// Creates Config Services
	err := RegisterSingleton(func() (c config.Config) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})

	// Creates Log Services
	err = RegisterSingleton(func(appConfig config.Config) logs.DataLogger {
		return logs.NewDefaultLog(appConfig)
	})
	if err != nil {
		panic(err)
	}
}
