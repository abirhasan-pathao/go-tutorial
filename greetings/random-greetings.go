package greetings

import "fmt"

func init() {
	println("randomgreetings package initialized")
}

func RandomGreet(name string) string {
	m := map[int]string{0: "Hi", 1: "Greetings", 2: "Salutations", 3: "Howdy", 4: "Hey there"}
	fmt.Println("RandomGreet called")
	return fmt.Sprintf("%s", m[len(name)%5]+", "+name+"!")
}
