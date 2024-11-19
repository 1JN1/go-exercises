package main

import (
	"fmt"
	"go-exercises/03/stack"
)

func main() {

	s := stack.NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Peek())

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}

}
