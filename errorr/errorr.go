// loop, controllflow, functions, variables

package main

import (
	"fmt"
	"math"
)

type NegativeSqrtError float64

func (e NegativeSqrtError) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt_newton_rufson(x float64) (float64, error) {
	if x < 0 {
		return 0, NegativeSqrtError(x)
	} else {
		precision := 0.001
		z := 1.0
		for math.Abs(z*z-x) >= precision {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt_newton_rufson(-8))
}
