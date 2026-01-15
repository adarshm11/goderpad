package services

import (
	"goderpad/models"
	"goderpad/utils"
)

// CreateRoom creates a new room with the given details.
// It does NOT add a user to the room, that is handled in the JoinRoom function.
func CreateRoom(userID, name, roomName string) (string, error) {
	roomID := utils.GenerateUUID()
	room := models.NewRoom(roomID, roomName)

	hub := models.GetHub()
	if err := hub.AddRoom(room); err != nil {
		return "", err
	}

	return roomID, nil
}

func JoinRoom(userID, name, roomID string) (map[string]any, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return nil, models.ErrRoomNotFound
	}

	user := models.CreateUser(userID, name)
	room.AddUser(user)

	response := map[string]any{
		"roomName": room.RoomName,
		"document": room.Document,
		"users":    room.GetCurrentUsers(),
	}
	return response, nil
}

func GetRoomName(roomID string) (string, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return "", models.ErrRoomNotFound
	}
	return room.RoomName, nil
}
