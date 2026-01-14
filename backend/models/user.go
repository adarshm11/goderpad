package models

type User struct {
	UserID         string
	Name           string
	CursorPosition CursorPosition
}

type CursorPosition struct {
	Line   int
	Column int
}

func CreateUser(userID, name string) *User {
	return &User{
		UserID:         userID,
		Name:           name,
		CursorPosition: CursorPosition{Line: 0, Column: 0},
	}
}
