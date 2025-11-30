package models

type RegisterRequest struct {
	User *User
	Room *Room
}

type UnregisterRequest struct {
	User *User
	Room *Room
}
