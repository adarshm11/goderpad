package services

import (
	"goderpad/models"
	"goderpad/utils"
)

func GetRoomInfo(roomID string) (map[string]any, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return map[string]any{
			"success": false,
			"error":   "Room not found",
		}, nil
	}
	return map[string]any{
		"success": true,
		"data": map[string]any{
			"roomName": room.RoomName,
		},
	}, nil
}

func CreateRoom(name, roomName string) (map[string]any, error) {
	roomID := utils.GenerateUUID()
	room := models.NewRoom(roomID, roomName, utils.GetCurrentTimestamp())
	hub := models.GetHub()
	hub.SetRoom(roomID, room)

	user := CreateUser(name)
	room.AddUser(user)

	return map[string]any{
		"success": true,
		"data": map[string]any{
			"roomId": room.RoomID,
			"userId": user.UserID,
		},
	}, nil
}

func JoinRoom(name, roomID string) (map[string]any, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomID)
	if !exists {
		return map[string]any{
			"success": false,
			"error":   "Room not found",
		}, nil
	}

	// Check for duplicate user
	if IsUserInRoom(roomID, name) {
		return map[string]any{
			"success": false,
			"error":   "User with this name already exists in the room",
		}, nil
	}

	user := CreateUser(name)
	room.AddUser(user)

	return map[string]any{
		"success": true,
		"data": map[string]any{
			"roomName": room.RoomName,
			"userId":   user.UserID,
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
