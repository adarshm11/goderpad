package models

// CreateRoomRequest represents a request to create a new room
type CreateRoomRequest struct {
	UserID   string `json:"userId"`
	RoomName string `json:"roomName"`
}

// DeleteRoomRequest represents a request to delete a room
type DeleteRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

// JoinRoomRequest represents a request to join a room
type JoinRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}

// LeaveRoomRequest represents a request to leave a room
type LeaveRoomRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
}
