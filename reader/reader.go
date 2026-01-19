package main

import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	r := MyReader{}
	i := 100
	b := make([]byte, 1)
	for i > 0 {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %q\n", n, err, b)
		i--
	}
}
