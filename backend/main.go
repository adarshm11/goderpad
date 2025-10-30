package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Cursor represents a cursor position in the editor
type Cursor struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// Message represents a WebSocket message
type Message struct {
	Type    string  `json:"type"` // "text-change", "cursor-update", "user-joined", "user-left"
	UserID  string  `json:"userId"`
	RoomID  string  `json:"roomId"`
	Content string  `json:"content,omitempty"`
	Cursor  *Cursor `json:"cursor,omitempty"`
}

// Client represents a single WebSocket connection
type Client struct {
	conn   *websocket.Conn
	roomID string
	userID string
	role   string // "interviewer", "candidate"
	send   chan []byte
	hub    *Hub
}

// Room manages a two-client collaborative session (interviewer/candidate)
type Room struct {
	id         string
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	document   string
	mu         sync.RWMutex
}

// Hub manages all rooms in the system
type Hub struct {
	rooms    map[string]*Room // roomID -> Room
	mu       sync.RWMutex
}

func main() {
	fmt.Println("goderpad main fn started")
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	})
	router.Run(":8080")
}
