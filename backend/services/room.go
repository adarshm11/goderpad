package services

import (
	"fmt"
	"maps"
	"sync"
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
		Deleted:  false,
		Lock:     sync.RWMutex{},
	}
	hub.Lock.Unlock()
	return roomID, nil
}

// DeleteRoom deletes a room and queues all its users for unregistration if the requesting user is the owner
func DeleteRoom(request models.DeleteRoomRequest) error {
	hub := GetHub()

	// Step 1: Acquire the Hub's lock to check room existence and ownership
	hub.Lock.Lock()
	room, roomExists := hub.Rooms[request.RoomID]
	if !roomExists {
		hub.Lock.Unlock()
		return fmt.Errorf("room %s does not exist", request.RoomID)
	}
	room.Lock.Lock()
	if room.Deleted {
		room.Lock.Unlock()
		hub.Lock.Unlock()
		return fmt.Errorf("room %s has already been deleted", request.RoomID)
	}
	if room.Owner != request.UserID {
		room.Lock.Unlock()
		hub.Lock.Unlock()
		return &models.PermissionError{Message: fmt.Sprintf("user %s does not have permission to delete room %s", request.UserID, request.RoomID)}
	}

	usersToUnregister := make([]*models.User, 0, len(room.Users))
	for _, user := range room.Users {
		usersToUnregister = append(usersToUnregister, user)
	}
	room.Deleted = true
	room.Lock.Unlock()

	delete(hub.Rooms, request.RoomID)
	hub.Lock.Unlock()

	for _, user := range usersToUnregister {
		hub.Unregister <- models.UnregisterRequest{User: user, Room: room}
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
	roomsSnapshot := make(map[string]*models.Room, len(hub.Rooms))
	maps.Copy(roomsSnapshot, hub.Rooms)
	hub.Lock.RUnlock()

	candidateRoomIDs := make([]string, 0)
	for roomID, room := range roomsSnapshot {
		room.Lock.RLock()
		shouldExpire := len(room.Users) == 0 && util.TimeSince(room.LastUsed) > util.WeekInSeconds && !room.Deleted
		room.Lock.RUnlock()
		if shouldExpire {
			candidateRoomIDs = append(candidateRoomIDs, roomID)
		}
	}

	for _, roomID := range candidateRoomIDs {
		hub.Lock.Lock()
		room, exists := hub.Rooms[roomID]
		if !exists {
			hub.Lock.Unlock()
			continue
		}
		room.Lock.Lock()
		if room.Deleted || len(room.Users) != 0 || util.TimeSince(room.LastUsed) <= util.WeekInSeconds {
			room.Lock.Unlock()
			hub.Lock.Unlock()
			continue
		}
		room.Deleted = true
		delete(hub.Rooms, roomID)
		room.Lock.Unlock()
		hub.Lock.Unlock()
	}
}
