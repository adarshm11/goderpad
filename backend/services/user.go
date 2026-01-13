package services

import (
	"goderpad/models"
)

// CreateUser creates a new user with a generated UUID
func CreateUser(userID, name string) models.User {
	return models.User{
		UserID: userID,
		Name:   name,
		Send:   make(chan models.BroadcastMessage, 256), // Buffered channel
	}
}

// GetUsersByRoom retrieves all users in a specific room
func GetUsersByRoom(roomId string) ([]models.User, error) {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomId)
	if !exists {
		return nil, nil
	}
	return room.GetUsers(), nil
}

// RemoveUserFromRoom removes a user from a room
func RemoveUserFromRoom(roomId, userId string) error {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomId)
	if !exists {
		return nil
	}
	room.RemoveUser(userId)
	return nil
}

// IsUserInRoom checks if a user with the given name exists in a room
func IsUserInRoom(roomId, userName string) bool {
	hub := models.GetHub()
	room, exists := hub.GetRoom(roomId)
	if !exists {
		return false
	}
	users := room.GetUsers()
	for _, user := range users {
		if user.Name == userName {
			return true
		}
	}
	return false
}
