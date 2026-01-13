package services

import (
	"goderpad/models"
	"goderpad/utils"
)

func CreateRoom(userID, name, roomName string) (map[string]any, error) {
	roomID := utils.GenerateUUID()
	room := &models.Room{
		RoomID:    roomID,
		RoomName:  roomName,
		Users:     []models.User{},
		CreatedAt: utils.GetCurrentTimestamp(),
		Document:  utils.DEFAULT_CODE,
		Broadcast: make(chan models.BroadcastMessage, 256), // Buffered channel
	}
	hub := models.GetHub()
	hub.SetRoom(roomID, room)

	room.StartBroadcaster()

	user := CreateUser(userID, name)
	room.AddUser(user)

	return map[string]any{
		"success": true,
		"data": map[string]any{
			"roomId": room.RoomID,
		},
	}, nil
}

func JoinRoom(userID, name, roomID string) (map[string]any, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return map[string]any{
			"success": false,
			"error":   "Room not found",
		}, nil
	}

	user := CreateUser(userID, name)
	room.AddUser(user)

	return map[string]any{
		"success": true,
		"data": map[string]any{
			"roomName": room.RoomName,
			"users":    room.GetUsers(),
			"document": room.GetDocument(),
		},
	}, nil
}

func LeaveRoom(userID, roomID string) error {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return nil
	}
	room.RemoveUser(userID)
	return nil
}

// HandleRoomBroadcasts listens for broadcast messages in a room and forwards them to all users except the sender
func HandleRoomBroadcasts(room *models.Room) {
	for {
		broadcastMessage := <-room.Broadcast
		userID := broadcastMessage.Payload.(map[string]any)["userId"].(string)
		for _, user := range room.GetUsers() {
			if user.UserID != userID {
				user.Send <- broadcastMessage
			}
		}
	}
}
