package logic

import (
	"fmt"
	"strings"
	"sync"
)

// 使⽤两个  goroutine 交替打印序列，⼀个  goroutine 打印数字， 另外⼀
// 个  goroutine 打印字⺟， 最终效果如下：

// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func AlternatePrint() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				number <- true
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
}
