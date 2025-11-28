package models

type PermissionError struct {
	Message string
}

func (e *PermissionError) Error() string {
	return e.Message
}
