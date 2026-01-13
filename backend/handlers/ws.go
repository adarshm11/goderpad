package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"goderpad/models"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:5173"
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["roomId"]
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "Missing userId parameter", http.StatusBadRequest)
		return
	}

	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	userExists := false
	users := room.GetUsers()
	for _, user := range users {
		if user.UserID == userID {
			userExists = true
			break
		}
	}
	if !userExists {
		http.Error(w, "User not found in room", http.StatusForbidden)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	go handleWebSocketConnection(conn, roomID, userID)
}

func handleWebSocketConnection(conn *websocket.Conn, roomID string, userID string) {
	defer conn.Close()

	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return
	}

	// Find the user
	var user *models.User
	for i, u := range room.GetUsers() {
		if u.UserID == userID {
			user = &room.GetUsers()[i]
			break
		}
	}
	if user == nil {
		return
	}

	// Initialize the Send channel
	user.Send = make(chan models.BroadcastMessage, 256)
	defer close(user.Send)

	// Goroutine to write messages from Send channel to WebSocket
	go func() {
		for message := range user.Send {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.WriteJSON(message); err != nil {
				log.Printf("Failed to write message to user %s: %v", userID, err)
				return
			}
		}
	}()

	// Read incoming messages from client
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		var broadcastMessage models.BroadcastMessage
		if err := json.Unmarshal(message, &broadcastMessage); err != nil {
			continue
		}

		broadcastMessage.UserID = userID
		room.Broadcast <- broadcastMessage
	}
}
