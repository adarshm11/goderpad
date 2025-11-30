package services

import (
	"fmt"

	"goderpad/models"
)

// JoinRoom queues a user to join a room
func JoinRoom(user *models.User, roomID string) error {
	hub := GetHub()
	currentRoom := user.GetRoom()
	if currentRoom != nil {
		if currentRoom.ID == roomID {
			return nil // User is already in the room
		}
		hub.Unregister <- models.UnregisterRequest{User: user, Room: currentRoom}
	}
	hub.Lock.RLock()
	room, roomExists := hub.Rooms[roomID]
	if !roomExists {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s does not exist", roomID)
	}
	hub.Lock.RUnlock()

	room.Lock.RLock()
	if room.Deleted {
		room.Lock.RUnlock()
		return fmt.Errorf("room %s has been deleted", roomID)
	}
	room.Lock.RUnlock()

	hub.Register <- models.RegisterRequest{User: user, Room: room}
	return nil
}

// LeaveRoom queues a user to leave a room
func LeaveRoom(user *models.User, roomID string) error {

	room := user.GetRoom()
	if room == nil {
		return fmt.Errorf("user %s is not in any room", user.ID)
	}
	if room.ID != roomID {
		return fmt.Errorf("user %s is not in room %s", user.ID, roomID)
	}

	hub := GetHub()
	hub.Lock.RLock()
	_, roomExists := hub.Rooms[roomID]
	if !roomExists {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s does not exist", roomID)
	}
	hub.Lock.RUnlock()

	room.Lock.RLock()
	if room.Deleted {
		room.Lock.RUnlock()
		return fmt.Errorf("room %s has been deleted", roomID)
	}
	room.Lock.RUnlock()
	hub.Unregister <- models.UnregisterRequest{User: user, Room: room}
	return nil
}
