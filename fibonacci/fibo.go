//closure

package main

import "fmt"

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		res := i
		i, j = j, i+j
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
