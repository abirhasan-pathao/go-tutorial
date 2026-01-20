package main

import (
	"fmt"
	"time"
)

func listNum(x int) {
	for i := 0; i <= x; i++ {
		time.Sleep(100 * time.Millisecond) // if we dont wait here, the goroutine may not get time to execute
		fmt.Println(i)
	}
}

func main() {
	go listNum(10)
	listNum(10) // this runs in the main goroutine to let the previous goroutine to execute before main goroutine exits
	fmt.Println("finished")

}
