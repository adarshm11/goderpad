package models

type TextChangeEvent struct {
	User *User
	// To-Do: decide how text changes are represented
}

type CursorUpdateEvent struct {
	User   *User
	OldPos CursorPosition
	NewPos CursorPosition
}

type UserJoinedEvent struct {
	User *User
	Room *Room
}

type UserLeftEvent struct {
	User *User
	Room *Room
}

var eventType = map[string]any{
	EventTextChange:   TextChangeEvent{},
	EventCursorUpdate: CursorUpdateEvent{},
	EventUserJoined:   UserJoinedEvent{},
	EventUserLeft:     UserLeftEvent{},
}
