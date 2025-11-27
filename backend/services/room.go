package services

import (
	"fmt"

	"goderpad/models"
	"goderpad/util"
)

func CreateRoom(request models.CreateRoomRequest) (string, error) {
	hub := GetHub()

	roomId := util.GenerateUniqueRoomID()
	hub.Rooms[roomId] = &models.Room{
		ID:    roomId,
		Name:  request.RoomName,
		Owner: request.UserID,
		Users: make(map[string]*models.User),
	}
	// users will be registered when they join the room
	return roomId, nil
}

func DeleteRoom(request models.DeleteRoomRequest) error {
	hub := GetHub()

	room, roomExists := hub.Rooms[request.RoomID]
	if !roomExists {
		return fmt.Errorf("room %s does not exist", request.RoomID)
	}
	if room.Owner != request.UserID {
		return &models.PermissionError{Message: fmt.Sprintf("user %s does not have permission to delete room %s", request.UserID, request.RoomID)}
	}

	for _, user := range room.Users {
		user.Room = nil
		hub.Unregister <- user
	}
	delete(hub.Rooms, request.RoomID)
	return nil
}
