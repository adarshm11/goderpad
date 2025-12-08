package services

import (
	"goderpad/models"
	"goderpad/utils"
)

func CreateRoom(request models.CreateRoomRequest) (string, error) {
	hub := GetHub()
	roomID := utils.GenerateUUID()

	room := &models.Room{
		ID:       roomID,
		RoomName: request.RoomName,
		Users:    make(map[string]*models.User),
		Owner:    &models.User{ID: request.UserID},
		Document: "", // empty document at creation
		LastUsed: utils.GetCurrentUnixTimestamp(),
	}

	// TODO: save the Room document to MongoDB

	// Owner won't be added right away -> only when they open the room (join), they will be added
	hub.Register <- room
	return roomID, nil
}

// DeleteRoom is called when there are no users left in a Room after it's done -> TODO: Verify this logic.
func DeleteRoom(request models.DeleteRoomRequest) {
	hub := GetHub()
	room := hub.GetRoom(request.RoomID)
	if room != nil {
		hub.Unregister <- room
	}
}
