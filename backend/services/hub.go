package services

import (
	"sync"

	"goderpad/models"
)

var hub = &models.Hub{
	Rooms:      make(map[string]*models.Room),
	Register:   make(chan *models.User),
	Unregister: make(chan *models.User),
	Broadcast:  make(chan models.Event),
	Lock:       sync.RWMutex{},
}

var stopChan = make(chan struct{})
var stopOnce sync.Once

func RegisterUsers() {
	for {
		select {
		case user := <-hub.Register:
			room := user.GetRoom()
			if room == nil {
				continue
			}
			room.Lock.Lock()
			room.Users[user.ID] = user
			room.Lock.Unlock()
		case <-stopChan:
			return
		}
	}
}

func UnregisterUsers() {
	for {
		select {
		case user := <-hub.Unregister:
			room := user.GetRoom()
			if room == nil {
				continue
			}
			room.Lock.Lock()
			delete(room.Users, user.ID)
			room.Lock.Unlock()
			user.Room = nil
		case <-stopChan:
			return
		}
	}
}

func GetHub() *models.Hub {
	return hub
}

func StopHub() {
	stopOnce.Do(func() {
		close(stopChan)
		expireTicker.Stop()
	})
}
