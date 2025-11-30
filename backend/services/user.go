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
	if !roomExists {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s does not exist", roomID)
	}

	user.SetRoom(room)
	hub.Lock.RUnlock()

	hub.Register <- user
	return nil
}

// LeaveRoom queues a user to leave a room
func LeaveRoom(user *models.User, roomID string) error {
	hub := GetHub()
	hub.Lock.Lock()
	room, roomExists := hub.Rooms[roomID]
	if !roomExists {
		hub.Lock.Unlock()
		return fmt.Errorf("room %s does not exist", roomID)
	}

	userRoom := user.GetRoom()
	if userRoom == nil {
		hub.Lock.Unlock()
		return fmt.Errorf("user %s is not in any room", user.ID)
	}
	if userRoom.ID != room.ID {
		hub.Lock.Unlock()
		return fmt.Errorf("user %s is not in room %s", user.ID, roomID)
	}
	hub.Lock.Unlock()
	hub.Unregister <- user
	return nil
}
