package util

import (
	"github.com/google/uuid"
)

func GenerateUniqueRoomID() string {
	return uuid.New().String()
}
