package main

import (
	"go-web-platform/services"
	"go-web-platform/view"
)

func main() {
	services.RegisterDefaultServices()
	view.Start()
}
