package app

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client struct for websocket connection and message sending
type Client struct {
	ID   string
	Conn *websocket.Conn
}

// Hub is a struct that holds all the clients and the messages that are sent to them
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool
	//Unregistered clients.
	unregister chan *Client
	// Register requests from the clients.
	register chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		unregister: make(chan *Client),
		register:   make(chan *Client),
	}
}



// Core function to run the hub
func (h *Hub) Run() {
	for {
		select {
		// Register a client.
		case client := <-h.register:
			h.RegisterNewClient(client)
			// Unregister a client.
		case client := <-h.unregister:
			h.RemoveClient(client)

		}
	}
}

// function check if room exists and if not create it and add client to it
func (h *Hub) RegisterNewClient(client *Client) {
	connections := h.clients[client.ID]
	if connections == nil {
		connections = make(map[*Client]bool)
		h.clients[client.ID] = connections
	}
	h.clients[client.ID][client] = true

	fmt.Println("Size of clients: ", len(h.clients[client.ID]))
}

// function to remvoe client from room
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.clients[client.ID]; ok {
		delete(h.clients[client.ID], client)
		fmt.Println("Removed client")
	}
}
