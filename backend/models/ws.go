package models

type MessageType string

const (
	UserJoinedMessageType   MessageType = "user_joined"
	UserLeftMessageType     MessageType = "user_left"
	CursorUpdateMessageType MessageType = "cursor_update"
	CodeUpdateMessageType   MessageType = "code_update"
)

type BroadcastMessage struct {
	UserID  string `json:"userId"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
