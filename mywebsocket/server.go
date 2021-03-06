package mywebsocket

import (
	"golang.org/x/net/websocket"
)

type Server struct {
	clientCount int
	clients     map[int]*Client
	// channels
	addClientCh    chan *Client
	removeClientCh chan *Client
	messageCh      chan *Message
}

func NewServer() *Server {
	return &Server{
		clientCount:    0,
		clients:        map[int]*Client{},
		addClientCh:    make(chan *Client),
		removeClientCh: make(chan *Client),
		messageCh:      make(chan *Message),
	}
}

func (server *Server) addClient(client *Client) {
	server.clientCount++
	client.Id = server.clientCount
	server.clients[client.Id] = client
}

func (server *Server) removeClient(client *Client) {
	delete(server.clients, client.Id)
}

func (server *Server) sendMessage(message *Message) {
	for _, client := range server.clients {
		if client.Room != message.Room {
			continue
		}
		c := client
		go func() { c.Send(message) }()
	}
}

func (server *Server) Start() {
	for {
		select {
		case client := <-server.addClientCh:
			server.addClient(client)
		case client := <-server.removeClientCh:
			server.removeClient(client)
		case message := <-server.messageCh:
			server.sendMessage(message)
		}
	}
}

func (server *Server) WebsocketHandler(room string) websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		client := NewClient(room, ws, server.removeClientCh, server.messageCh)
		server.addClientCh <- client
		client.Start()
	})
}
