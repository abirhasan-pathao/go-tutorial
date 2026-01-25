package greetings

import "fmt"

func init() {
	println("greetings package initialized")
}

func Hello(name string) string {
	fmt.Println("Hello function called")
	return "Hello, " + name + "!"
}
