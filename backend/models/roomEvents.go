package models

type CreateRoomRequest struct {
	RoomID string `json:"roomId"`
}

type DeleteRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type JoinRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

type LeaveRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}
