package sync

import (
	"fmt"
	"sync"
)

func SyncMap() {
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")

	val, ok := m.Load("cat")
	fmt.Println("SyncMap:")
	if ok {
		fmt.Println(len(val.(string)))
	}
}
