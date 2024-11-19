package main

import (
	"fmt"
	"go-exercises/03/stack"
	"strconv"
)

func main() {

	s := stack.NewStackWithCap[string](1000)

	for i := 0; i < 1000; i++ {
		s.Push("第" + strconv.Itoa(i+1) + "个元素")
	}

	fmt.Println(s.Peek())

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}

}
