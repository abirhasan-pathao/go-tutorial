package main

import "golang.org/x/tour/tree"

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func same(ch1, ch2 chan int) bool {
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		}
		if ok1 != ok2 || v1 != v2 {
			return false
		}
	}

}
func main() {
	tree1 := tree.New(5)
	tree2 := tree.New(3)

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		walk(tree1, ch1)
		close(ch1)
	}()
	go func() {
		walk(tree2, ch2)
		close(ch2)
	}()

	println(same(ch1, ch2))
}
