package services

import (
	"sync"

	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.User, 100),
	Unregister: make(chan *models.User, 100),
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
		case user := <-hub.Register:
			hub.Lock.RLock()
			room := user.GetRoom()
			hub.Lock.RUnlock()
			if room == nil {
				continue
			}
			room.UpdateLastUsed()
			room.Lock.Lock()
			room.Users[user.ID] = user
			room.Lock.Unlock()
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
		case user := <-hub.Unregister:
			hub.Lock.RLock()
			room := user.GetRoom()
			if room == nil {
				hub.Lock.RUnlock()
				continue
			}
			if _, ok := hub.Rooms[room.ID]; !ok {
				hub.Lock.RUnlock()
				continue
			}
			room.Lock.Lock()
			hub.Lock.RUnlock()
			delete(room.Users, user.ID)
			room.UpdateLastUsed()
			room.Lock.Unlock()
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
