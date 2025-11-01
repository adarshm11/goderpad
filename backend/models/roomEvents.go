package models

type CreateRoomRequest struct {
	UserID string `json:"userId"`
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
