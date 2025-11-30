package services

import (
	"fmt"
	"maps"
	"time"

	"goderpad/models"
	"goderpad/util"
)

var expireTicker = time.NewTicker(time.Hour)

// CreateRoom creates a new room and returns its ID
func CreateRoom(request models.CreateRoomRequest) (string, error) {
	hub := GetHub()

	roomID := util.GenerateUniqueRoomID()
	hub.Lock.Lock()
	hub.Rooms[roomID] = &models.Room{
		ID:       roomID,
		Name:     request.RoomName,
		Owner:    request.UserID,
		Users:    make(map[string]*models.User),
		LastUsed: time.Now(),
	}
	hub.Lock.Unlock()
	return roomID, nil
}

// DeleteRoom deletes a room and queues all its users for unregistration if the requesting user is the owner
func DeleteRoom(request models.DeleteRoomRequest) error {
	hub := GetHub()

	hub.Lock.Lock()
	room, roomExists := hub.Rooms[request.RoomID]
	if !roomExists {
		hub.Lock.Unlock()
		return fmt.Errorf("room %s does not exist", request.RoomID)
	}
	if room.Owner != request.UserID {
		hub.Lock.Unlock()
		return &models.PermissionError{Message: fmt.Sprintf("user %s does not have permission to delete room %s", request.UserID, request.RoomID)}
	}

	room.Lock.Lock()
	usersToUnregister := make([]*models.User, 0, len(room.Users))
	for _, user := range room.Users {
		usersToUnregister = append(usersToUnregister, user)
	}
	room.Lock.Unlock()

	delete(hub.Rooms, request.RoomID)
	hub.Lock.Unlock()

	for _, user := range usersToUnregister {
		hub.Unregister <- user
	}
	return nil
}

// StartRoomExpiration starts a background process that periodically expires inactive rooms
// This is a goroutine that is started in main.go
func StartRoomExpiration() {
	for range expireTicker.C {
		ExpireRooms()
	}
}

// ExpireRooms removes rooms that have been inactive for more than a week and have no users
func ExpireRooms() {
	hub := GetHub()

	hub.Lock.RLock()
	roomsCopy := make(map[string]*models.Room)
	maps.Copy(roomsCopy, hub.Rooms)
	hub.Lock.RUnlock()

	var roomsToDelete []string

	for roomID, room := range roomsCopy {
		room.Lock.Lock()
		lastUsed := room.LastUsed
		numUsers := len(room.Users)
		room.Lock.Unlock()
		if numUsers == 0 && util.TimeSince(lastUsed) > util.WeekInSeconds {
			roomsToDelete = append(roomsToDelete, roomID)
		}
	}

	for _, roomID := range roomsToDelete {
		hub.Lock.Lock()
		room, exists := hub.Rooms[roomID]
		if !exists {
			hub.Lock.Unlock()
			continue
		}
		room.Lock.Lock()
		numUsers := len(room.Users)
		lastUsed := room.LastUsed
		if numUsers == 0 && util.TimeSince(lastUsed) > util.WeekInSeconds {
			delete(hub.Rooms, roomID)
		}
		room.Lock.Unlock()
		hub.Lock.Unlock()
	}
}
