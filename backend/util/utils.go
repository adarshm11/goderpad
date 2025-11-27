package util

import (
	"github.com/google/uuid"

	"time"
)

func GenerateUniqueRoomID() string {
	return uuid.New().String()
}

func TimeSince(t time.Time) int64 {
	return int64(time.Since(t).Seconds())
}

const WeekInSeconds = 7 * 24 * 60 * 60
