package main

import (
	"github.com/dfree1645/goa-websocket-study/app"
	"github.com/dfree1645/goa-websocket-study/mywebsocket"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

// WebsocketController implements the websocket resource.
type WebsocketController struct {
	*goa.Controller
	websocketCtl *mywebsocket.Server
}

// NewWebsocketController creates a websocket controller.
func NewWebsocketController(service *goa.Service) *WebsocketController {
	websockObj := mywebsocket.NewServer()
	go websockObj.Start()

	return &WebsocketController{
		Controller:   service.NewController("WebsocketController"),
		websocketCtl: websockObj,
	}
}

// Connect runs the connect action.
func (c *WebsocketController) Connect(ctx *app.ConnectWebsocketContext) error {
	c.ConnectWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// ConnectWSHandler establishes a websocket connection to run the connect action.
func (c *WebsocketController) ConnectWSHandler(ctx *app.ConnectWebsocketContext) websocket.Handler {
	return c.websocketCtl.WebsocketHandler(ctx.Room)

}
