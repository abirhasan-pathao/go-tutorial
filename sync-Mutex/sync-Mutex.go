package main

import (
	"sync"
	"time"
)

func increment(mutex *sync.Mutex, counter *int) {
	mutex.Lock() // without the locks we wont get the expected result due to race conditions
	*counter++
	mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go increment(&mutex, &counter)
	}

	// Wait for goroutines to finish (in a real program, use sync.WaitGroup)
	time.Sleep(100 * time.Millisecond)
	println("Final Counter:", counter)

}
