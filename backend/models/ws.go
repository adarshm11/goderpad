package models

type BroadcastMessage struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
