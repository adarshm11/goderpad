package models

type CreateRoomRequest struct {
	UserID   string `json:"userId"`
	RoomName string `json:"roomName"`
}

type DeleteRoomRequest struct {
	UserID string `json:"userId"`
	RoomID string `json:"roomId"`
}

type JoinRoomRequest struct {
	UserID string `json:"userId"`
	RoomID string `json:"roomId"`
}

type LeaveRoomRequest struct {
	UserID string `json:"userId"`
	RoomID string `json:"roomId"`
}
