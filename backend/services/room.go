package services

import (
	"fmt"
	"time"

	"goderpad/models"
	"goderpad/util"
)

var expireTicker = time.NewTicker(time.Hour)

func CreateRoom(request models.CreateRoomRequest) (string, error) {
	hub := GetHub()

	roomId := util.GenerateUniqueRoomID()
	hub.Lock.Lock()
	hub.Rooms[roomId] = &models.Room{
		ID:       roomId,
		Name:     request.RoomName,
		Owner:    request.UserID,
		Users:    make(map[string]*models.User),
		LastUsed: time.Now(),
	}
	hub.Lock.Unlock()
	return roomId, nil
}

func DeleteRoom(request models.DeleteRoomRequest) error {
	hub := GetHub()

	hub.Lock.Lock()
	defer hub.Lock.Unlock()
	room, roomExists := hub.Rooms[request.RoomID]
	if !roomExists {
		return fmt.Errorf("room %s does not exist", request.RoomID)
	}
	if room.Owner != request.UserID {
		return &models.PermissionError{Message: fmt.Sprintf("user %s does not have permission to delete room %s", request.UserID, request.RoomID)}
	}

	room.Lock.Lock()
	for _, user := range room.Users {
		hub.Unregister <- user
	}
	room.Lock.Unlock()

	delete(hub.Rooms, request.RoomID)
	return nil
}

func StartRoomExpiration() {
	for range expireTicker.C {
		ExpireRooms()
	}
}

func ExpireRooms() {
	hub := GetHub()
	hub.Lock.Lock()
	defer hub.Lock.Unlock()

	for roomId, room := range hub.Rooms {
		room.Lock.Lock()
		lastUsed := room.LastUsed
		numUsers := len(room.Users)
		room.Lock.Unlock()
		if numUsers == 0 && util.TimeSince(lastUsed) > util.WeekInSeconds {
			delete(hub.Rooms, roomId)
		}
	}
}
