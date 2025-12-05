package services

import (
	"errors"

	"goderpad/models"
)

// JoinRoom takes in a UserID and RoomID, adds the User to the Room's Users map.
// There may be a case where the Room's document exists in the database but the Room is not
// currently active in memory. In that case, we create the Room in memory first.
func JoinRoom(roomID string, user *models.User) error {
	hub := GetHub()
	room := hub.GetRoom(roomID)
	if room == nil {
		// TODO: Check if the Room exists in database -> if it does, recreate the Room in memory
		// If not, return 404 error
		return errors.New("room not found")
	}

	room.AddUser(user)
	return nil
}

// LeaveRoom takes in a RoomID and a UserID and removes the User from the Room's Users map.
// If the Room becomes empty after this operation, it unregisters the Room.
func LeaveRoom(roomID string, userID string) error {
	hub := GetHub()
	room := hub.GetRoom(roomID)
	if room == nil {
		return errors.New("room not found")
	}

	if room.RemoveUser(userID) {
		hub.Unregister <- room
	}
	return nil
}
