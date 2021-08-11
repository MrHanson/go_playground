package sync

import (
	"fmt"
	"sync"
)

var once sync.Once

func PrintOnce() {
	once.Do(func() {
		fmt.Printf("print only once")
	})
	once.Do(func() {
		fmt.Printf("print second func")
	})
}

func OnceRun() {
	// call three times
	PrintOnce()
	PrintOnce()
	PrintOnce()
}
