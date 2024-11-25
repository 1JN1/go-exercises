package main

import "fmt"

func main() {

	add := func() func(x, y int) int {

		// 记录内层匿名函数执行次数
		count := 0
		return func(x, y int) int {

			count++
			fmt.Printf("已经执行了 %d 次\n", count)

			return x + y
		}

	}()

	x, y := 1, 1

	for i := 0; i < 100; i++ {
		x = add(x, y)
	}

	fmt.Println(x, y)

}
