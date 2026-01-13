package models

type User struct {
	UserID         string                `json:"userId"`
	Name           string                `json:"userName"`
	CursorPosition *CursorPosition       `json:"cursorPosition"`
	Send           chan BroadcastMessage `json:"-"`
}

type CursorPosition struct {
	LineNumber int `json:"lineNumber"`
	Column     int `json:"column"`
}
