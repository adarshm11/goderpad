package services

import (
	"fmt"

	"goderpad/models"
)

// JoinRoom queues a user to join a room
func JoinRoom(user *models.User, roomID string) error {
	hub := GetHub()
	hub.Lock.RLock()
	room, roomExists := hub.Rooms[roomID]
	hub.Lock.RUnlock()
	if !roomExists {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	user.SetRoom(room)
	hub.Register <- user
	return nil
}

// LeaveRoom queues a user to leave a room
func LeaveRoom(user *models.User, roomID string) error {
	hub := GetHub()
	hub.Lock.RLock()
	room, roomExists := hub.Rooms[roomID]
	hub.Lock.RUnlock()
	if !roomExists {
		return fmt.Errorf("room %s does not exist", roomID)
	}

	userRoom := user.GetRoom()
	if userRoom == nil {
		return fmt.Errorf("user %s is not in any room", user.ID)
	}
	if userRoom.ID != room.ID {
		return fmt.Errorf("user %s is not in room %s", user.ID, roomID)
	}

	hub.Unregister <- user
	return nil
}
