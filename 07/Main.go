package main

import "fmt"

func main() {

	nums := []int{3, 1, 1, 3}

	fmt.Println(nums)

	deleteAllElem(&nums, 1)

	fmt.Println(nums)

}

// deleteAllElem 从切片中删除所有指定的元素elem
func deleteAllElem[T int | float32 | float64 | string](nums *[]T, elem T) {

	for i := 0; i < len(*nums); i++ {

		if (*nums)[i] == elem {

			*nums = append((*nums)[:i], (*nums)[i+1:]...)
			i--
		}

	}
}
