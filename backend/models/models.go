package models

import "github.com/gorilla/websocket"

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Event
}

type Room struct {
	ID      string
	Clients map[string]*Client
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Hub  *Hub
}

type Event struct {
	// Possible events: text-change, cursor-update, user-joined, user-left, etc

}
