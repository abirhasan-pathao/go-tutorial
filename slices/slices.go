// a slice is a dynamically-sized, flexible view into the elements of an array.
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// a slice is a vector that gives access to the underlying array

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		ans[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			ans[i][j] = uint8(i + j)
		}
	}
	return ans
}

func main() {
	pic.Show(Pic) // display the slice as an image
}

// copy the code and replace the code here https://go.dev/tour/moretypes/18 to see the actual image
