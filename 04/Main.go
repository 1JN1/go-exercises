package main

import (
	"fmt"
	"go-exercises/04/queue"
)

func main() {

	q := queue.Queue[int]{}

	for i := 0; i < 1000000; i++ {
		q.EnQueue(i + 1)
	}

	for !q.IsEmpty() {
		fmt.Println(q.DeQueue())
	}

}
