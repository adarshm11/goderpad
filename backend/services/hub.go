package services

import (
	"sync"
	"time"

	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan models.RegisterRequest, 100),
	Unregister: make(chan models.UnregisterRequest, 100),
	Broadcast:  make(chan models.Event),
	Lock:       sync.RWMutex{},
}

var stopChan = make(chan struct{})
var stopOnce sync.Once

// RegisterUsers processes user registration requests from the Register channel
// This is a goroutine that is started in main.go
func RegisterUsers() {
	for {
		select {
		case request := <-hub.Register:
			hub.Lock.RLock()
			user, room := request.User, request.Room
			hub.Lock.RUnlock()
			if room == nil {
				continue
			}
			room.Lock.Lock()
			if room.Deleted {
				room.Lock.Unlock()
				continue
			}
			room.LastUsed = time.Now()
			room.Users[user.ID] = user
			room.Lock.Unlock()

			user.SetRoom(room)
		case <-stopChan:
			return
		}
	}
}

// UnregisterUsers processes user unregistration requests from the Unregister channel
// This is a goroutine that is started in main.go
func UnregisterUsers() {
	for {
		select {
		case request := <-hub.Unregister:
			hub.Lock.RLock()
			user, room := request.User, request.Room
			if room == nil {
				hub.Lock.RUnlock()
				continue
			}
			// check if Room exists in the hub
			if _, ok := hub.Rooms[room.ID]; !ok && !room.Deleted {
				hub.Lock.RUnlock()
				continue
			}
			hub.Lock.RUnlock()
			room.RemoveUser(user.ID)
			room.UpdateLastUsed()
			user.SetRoom(nil)
		case <-stopChan:
			return
		}
	}
}

// GetHub returns the singleton Hub instance
func GetHub() *models.Hub {
	return hub
}

// StopHub gracefully shuts down the hub by closing the stop channel and stopping the expire ticker
func StopHub() {
	stopOnce.Do(func() {
		close(stopChan)
		expireTicker.Stop()
	})
}
