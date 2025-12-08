package utils

import (
	"time"
)

func GetCurrentUnixTimestamp() int64 {
	return time.Now().Unix()
}

func CreateTicker(duration time.Duration) *time.Ticker {
	return time.NewTicker(duration)
}

const HOUR = 1 * time.Hour
