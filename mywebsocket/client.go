package mywebsocket

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Client struct {
	Id   int
	Room string
	ws   *websocket.Conn
	// share with server
	removeClientCh chan *Client
	messageCh      chan *Message
}

func NewClient(room string, ws *websocket.Conn, remove chan *Client, message chan *Message) *Client {
	return &Client{
		Room:           room,
		ws:             ws,
		removeClientCh: remove,
		messageCh:      message,
	}
}

func (client *Client) Start() {
	for {
		var message string
		err := websocket.Message.Receive(client.ws, &message)
		if err != nil {
			client.removeClientCh <- client
			return
		} else {
			client.messageCh <- &Message{Room: client.Room, Body: message}
		}
	}
}

func (client *Client) Send(message *Message) {
	err := websocket.JSON.Send(client.ws, message)
	if err != nil {
		fmt.Println(message)
	}
}

func (client *Client) Close() {
	client.ws.Close()
}
