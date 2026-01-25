// defer, stacking behavior of defer

package main

import "fmt"

func deferedPrint(n int) int {
	fmt.Println(n)
	defer func() {
		n++
		fmt.Println("deferred:", n)
	}()
	return n // value returned before deferred function is executed
}

func main() {

	fmt.Println("counting")

	defer fmt.Println("hello")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // deferred calls are stacked
	}

	fmt.Println("done")
	defer fmt.Println("this is the last line to be executed")

	result := deferedPrint(5)
	fmt.Println("result:", result)
}
