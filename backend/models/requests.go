package models

type CreateRoomRequest struct {
	Name     string `json:"name"`
	RoomName string `json:"roomName"`
}

type JoinRoomRequest struct {
	Name   string `json:"name"`
	RoomId string `json:"roomId"`
}
