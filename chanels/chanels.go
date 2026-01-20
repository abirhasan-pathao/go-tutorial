//chanels are synchronised by default.
//unbuffered chanels block the sender until the receiver receives the value and vice versa.
//buffered chanels allow sending multiple values without blocking until the buffer is full.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(time.Duration(rand.Int()%10) * time.Millisecond) // simulate variable computation time
	c <- sum                                                    // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
