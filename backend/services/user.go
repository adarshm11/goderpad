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
	if room.Deleted {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s has been deleted", roomID)
	}

	hub.Lock.RUnlock()

	hub.Register <- models.RegisterRequest{User: user, Room: room}
	return nil
}

// LeaveRoom queues a user to leave a room
func LeaveRoom(user *models.User, roomID string) error {
	hub := GetHub()
	hub.Lock.RLock()
	room, roomExists := hub.Rooms[roomID]
	if !roomExists {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s does not exist", roomID)
	}
	if room.Deleted {
		hub.Lock.RUnlock()
		return fmt.Errorf("room %s has been deleted", roomID)
	}
	// check if the user is actually in the room
	userRoom := user.GetRoom()
	if userRoom == nil {
		hub.Lock.RUnlock()
		return fmt.Errorf("user %s is not in any room", user.ID)
	}
	// check if the user is leaving the specified room
	if userRoom.ID != room.ID {
		hub.Lock.RUnlock()
		return fmt.Errorf("user %s is not in room %s", user.ID, roomID)
	}
	hub.Lock.RUnlock()
	hub.Unregister <- models.UnregisterRequest{User: user, Room: room}
	return nil
}
