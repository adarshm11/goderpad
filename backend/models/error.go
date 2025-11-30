package models

// PermissionError represents an error due to insufficient permissions
type PermissionError struct {
	Message string
}

func (e *PermissionError) Error() string {
	return e.Message
}
