package services

import (
	"fmt"

	"goderpad/models"
)

func JoinRoom(user models.User, roomId string) error {
	hub := GetHub()
	room, roomExists := hub.Rooms[roomId]
	if !roomExists {
		return fmt.Errorf("room %s does not exist", roomId)
	}

	user.Room = room
	hub.Register <- &user
	return nil
}

func LeaveRoom(user models.User, roomId string) error {
	hub := GetHub()
	room, roomExists := hub.Rooms[roomId]
	if !roomExists {
		return fmt.Errorf("room %s does not exist", roomId)
	}

	if user.Room.ID != room.ID {
		return fmt.Errorf("user %s is not in room %s", user.ID, roomId)
	}

	hub.Unregister <- &user
	return nil
}
