// map, ranged loops, strings.Fields

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	st := strings.Fields(s)
	for _, v := range st {
		m[v]++
	}
	return m
}

func main() {
	fmt.Println(WordCount("a bad brown fox is very bad fox"))
}
