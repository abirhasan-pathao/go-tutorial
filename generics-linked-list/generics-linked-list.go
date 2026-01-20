//generics, linked list, pointer

package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func Insert[T any](head **List[T], val T) {
	if *head == nil {
		*head = &List[T]{val: val}
		return
	}
	curr := *head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &List[T]{val: val}
}

func (l *List[T]) Print() {
	fmt.Printf("%T List: ", l.val)
	for l != nil {
		fmt.Print(l.val)
		l = l.next
		if l != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func main() {
	var headint *List[int]

	Insert(&headint, 10)
	Insert(&headint, 20)
	Insert(&headint, 30)
	headint.Print()

	var headstring *List[string]

	Insert(&headstring, "hello")
	Insert(&headstring, "world")
	Insert(&headstring, "generics")
	headstring.Print()

}
