package sync

import "sync"

var mutex sync.Mutex
var rwMutex sync.RWMutex

func MutexFailed() {
	mutex.Lock()
	defer mutex.Unlock()

	mutex.Lock()
	defer mutex.Unlock()
}

func RwMutexFailed() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	rwMutex.Lock()
	defer rwMutex.Unlock()
}
