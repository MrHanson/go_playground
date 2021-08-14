package channel

import (
	"fmt"
	"time"
)

func RangeRun() {
	go func() {
		fmt.Println("First gorutine")
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		fmt.Println("Second gorutine")
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}
