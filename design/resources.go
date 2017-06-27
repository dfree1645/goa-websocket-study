package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/public/*filepath", "public/")
})

var _ = Resource("websocket", func() {
	Action("connect", func() {
		Routing(GET("ws"))
		Scheme("ws")
		Description("websocket server")
		Params(func() {
			Param("room", String, "ルームID チャットルーム的な")
			Required("room")
		})
		Response(SwitchingProtocols)
	})

})
