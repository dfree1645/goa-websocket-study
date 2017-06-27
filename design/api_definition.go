package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("websocket test", func() {
	Title("Websocket test")
	Description("websocket試し")
	Contact(func() {
		Name("dfree1645")
		Email("d.free1645@gmail.com")
		URL("https://github.com/dfree1645")
	})
	License(func() {
		Name("MIT")
		URL("")
	})
	Docs(func() {
		Description("wiki")
		URL("")
	})
	Host("localhost:8080")
	Scheme("http")
	//Scheme("https")
	BasePath("/api/v1")

	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		Headers("Content-Type", "X-Authorization")
		MaxAge(600)
		Credentials()
	})
})
