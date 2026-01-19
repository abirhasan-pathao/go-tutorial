// structs are typed collections of fields.

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // values assigned in order
	v2 = Vertex{X: 1}  // named field ommitted field gets default value
	v3 = Vertex{}      // all fields get default value
	p  = &Vertex{1, 2} // pointer of type Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
