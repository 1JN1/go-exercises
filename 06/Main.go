package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	n := 1000000000
	// 假设我们有一个数字数组
	numbers := make([]int, 0, n)

	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}

	startTime := time.Now()

	totalSum := sumSquare(numbers)

	endTime := time.Now()

	fmt.Println("总和是:", totalSum)
	fmt.Printf("计算%d个数的平法和耗费的时间为：%v\n", n, endTime.Sub(startTime))
}

// sumSquare 多线程计算平方和
func sumSquare[T int | float32 | float64](numbers []T) T {

	// 开启多少个Goroutine
	size := int(math.Sqrt(float64(len(numbers))))

	// 通道存放每一部分的平法和，向上取整
	sumChan := make(chan T, (len(numbers)+size-1)/size)

	var wg sync.WaitGroup

	for i := 0; i < len(numbers); i += size {

		wg.Add(1)

		end := int(math.Min(float64(i+size), float64(len(numbers))))

		go func(start, end int) {

			defer wg.Done()

			partialSum := T(0)

			for _, num := range numbers[start:end] {
				partialSum += num * num
			}

			// 将片段的和存入通道
			sumChan <- partialSum

		}(i, end)
	}

	go func() {
		wg.Wait()
		// 关闭通道
		close(sumChan)
	}()

	totalSum := T(0)

	for sum := range sumChan {
		totalSum += sum
	}

	return totalSum
}
