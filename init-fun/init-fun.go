// init functions run before main function
// multiple init functions are allowed and they run in the order they are defined in the file.
// init functions in imported packages run before the init functions in the main package.

package main

import (
	"fmt"
	"go-tutorial/greetings"
	"go-tutorial/greetings/long"
)

func init() {
	fmt.Println("init1 function called")
}

func init() {
	fmt.Println("init2 function called")
}

func main() {
	fmt.Println("main")
	fmt.Println(greetings.Hello("Alice"))
	fmt.Println(greetings.RandomGreet("Abir"))
	fmt.Println(long.LongGreeting("Bob"))
}

func init() {
	fmt.Println("init3 function called")
}
