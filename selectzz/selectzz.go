//select is like a switch but for channels. It allows a goroutine to wait on multiple communication operations.
//it will block until any one of its cases can run, then it executes that case. If multiple are ready, one is chosen at random.

package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("waiting...")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			time.Sleep(1 * time.Nanosecond)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
