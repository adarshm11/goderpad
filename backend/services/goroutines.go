package services

import (
	"context"
	"log"

	"goderpad/utils"
)

func RegisterRooms(ctx context.Context) {
	hub := GetHub()
	for {
		select {
		case room := <-hub.Register:
			hub.AddRoom(room)
		case <-ctx.Done():
			return
		}
	}
}

func UnregisterRooms(ctx context.Context) {
	hub := GetHub()
	for {
		select {
		case room := <-hub.Unregister:
			hub.RemoveRoom(room.ID)
		case <-ctx.Done():
			return
		}
	}
}

func ExpireRooms(ctx context.Context) {
	ticker := utils.CreateTicker(utils.HOUR)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check for expired rooms in DB
		case <-ctx.Done():
			log.Println("ExpireRooms goroutine exiting")
			return
		}
	}
}
