// loop, controllflow, functions, variables

package main

import (
	"fmt"
	"math"
)

func Sqrt_newton_rufson(x float64) float64 {
	precision := 0.001
	z := 1.0
	for math.Abs(z*z-x) >= precision {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Sqrt_binary_search(x float64) float64 {
	precision := 0.001
	low := 0.0
	high := x
	if x < 1 {
		high = 1
	}
	z := (low + high) / 2
	for math.Abs(z*z-x) >= precision {
		if z*z < x {
			low = z
		} else {
			high = z
		}
		z = (low + high) / 2
	}
	return z
}

func main() {
	fmt.Println(Sqrt_binary_search(8))
}
