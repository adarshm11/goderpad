package handlers

import (
	"goderpad/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Plan:
// Websocket handler listens for incoming connection requests in the form: /ws/:roomId
// Then it upgrades the connection for that user to a websocket connection
// Then the websocket handler listens for messages, extracts its roomId and userId
// Then it fetches the room from the hub using roomId and passes the message to the room's Broadcast channel
// The Room listens for messages on its Broadcast channel and funnels them to the respective users' Send channels
// That excludes the user who sent the message
// The users listen for messages on their Send channels and write them to their websocket connections
func WebSocketHandler(c *gin.Context) {
	log.Printf("WebSocket request received: %s", c.Request.URL.Path)
	roomID := c.Param("roomID")
	userID := c.Query("userId")
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	// The user has already "joined" the room via the HTTP endpoint
	// So we just need to set up their websocket connection
	user, userExists := room.CheckUserExists(userID)
	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found in room"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
	}

	user.Conn = conn

	// Start listening for messages from this user's websocket connection
	go readBroadcastsFromUser(user, room)
}

func readBroadcastsFromUser(user *models.User, room *models.Room) {
	defer closeUserConnection(user, room)

	for {
		var msg models.BroadcastMessage
		err := user.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		if msg.Type == "cursor_update" {
			if line, ok := msg.Payload["line"].(int); ok {
				if column, ok := msg.Payload["column"].(int); ok {
					user.UpdateCursorPosition(line, column)
				}
			}
		}
		room.Broadcast <- msg
	}
}

func closeUserConnection(user *models.User, room *models.Room) {
	user.Conn.Close()
	user.Close() // shut down the user's goroutines
	room.RemoveUser(user.UserID)
}
