// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "websocket test": websocket Resource Client
//
// Command:
// $ goagen
// --design=github.com/dfree1645/goa-websocket-study/design
// --out=$(GOPATH)\src\github.com\dfree1645\goa-websocket-study
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"golang.org/x/net/websocket"
	"net/url"
)

// ConnectWebsocketPath computes a request path to the connect action of websocket.
func ConnectWebsocketPath() string {

	return fmt.Sprintf("/api/v1/ws")
}

// websocket server
func (c *Client) ConnectWebsocket(ctx context.Context, path string, room string) (*websocket.Conn, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "ws"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("room", room)
	u.RawQuery = values.Encode()
	url_ := u.String()
	cfg, err := websocket.NewConfig(url_, url_)
	if err != nil {
		return nil, err
	}
	return websocket.DialConfig(cfg)
}