//go:generate goagen bootstrap -d github.com/dfree1645/goa-websocket-study/design

package main

import (
	"github.com/dfree1645/goa-websocket-study/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("websocket test")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "public" controller
	c := NewPublicController(service)
	app.MountPublicController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "websocket" controller
	c3 := NewWebsocketController(service)
	app.MountWebsocketController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
