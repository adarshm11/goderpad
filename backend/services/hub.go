package services

import (
	"log"
	"time"

	"goderpad/models"
)

var hub = models.GetHub()

func ExpireRooms() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		expiredCount := hub.ExpireRooms(60 * 60 * 24 * 7) // Expire rooms older than 1 week

		if expiredCount > 0 {
			log.Printf("Expired %d rooms\n", expiredCount)
		}
	}
}
