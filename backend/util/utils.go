package util

import (
	"time"

	"github.com/google/uuid"
)

// GenerateUniqueRoomID generates a unique room ID using UUID
func GenerateUniqueRoomID() string {
	return uuid.New().String()
}

// TimeSince returns the time elapsed in seconds since the given time
func TimeSince(t time.Time) int64 {
	return int64(time.Since(t).Seconds())
}

// WeekInSeconds represents the number of seconds in a week
const WeekInSeconds = 7 * 24 * 60 * 60
