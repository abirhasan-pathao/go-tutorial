// pointers

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Printf("p = %v, *p = %v, &p = %v, i = %v\n", p, *p, &p, i)
	*p = 21
	fmt.Printf("midified through p, now i = %v\n", i)

	p = &j
	*p = *p / 37
	fmt.Printf("p now points to j and can modify it so j = %v\n", j)
}
