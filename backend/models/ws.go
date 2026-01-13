package models

type MessageType string

const (
	MessageTypeUserJoined   MessageType = "user_joined"
	MessageTypeUserLeft     MessageType = "user_left"
	MessageTypeCursorUpdate MessageType = "cursor_update"
	MessageTypeCodeUpdate   MessageType = "code_update"
)

type BroadcastMessage struct {
	UserID  string      `json:"userId"`
	Type    MessageType `json:"type"`
	Payload any         `json:"payload"`
}

type UserJoinedEvent struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	RoomID string `json:"roomId"`
}

type UserLeftEvent struct {
	UserID string `json:"userId"`
	RoomID string `json:"roomId"`
}

type CursorUpdateEvent struct {
	UserID     string `json:"userId"`
	LineNumber int    `json:"lineNumber"`
	Column     int    `json:"column"`
}

type CodeUpdateEvent struct {
	UserID string `json:"userId"`
	Code   string `json:"code"`
}
