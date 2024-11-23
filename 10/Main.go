package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var num = 0

const MaxNum = 10

func main() {

	wg.Add(4)

	go producer("小明")
	go producer("小黄")
	go consumer("大黄")
	go consumer("大白")

	wg.Wait()
}

func producer(name string) {

	defer wg.Done()

	count := 0

	for {
		mutex.Lock()
		if num >= MaxNum {
			fmt.Printf("%s：仓库满了\n", name)
			mutex.Unlock()
		} else {
			// 模拟生成产品的时间
			time.Sleep(time.Second * 3)
			count++
			num++

			fmt.Printf(
				"%s：生产了一件产品，我已经生产了%d件产品，当前仓库剩余容量为 %d\n",
				name,
				count,
				MaxNum-num,
			)
			mutex.Unlock()
		}

	}

}

func consumer(name string) {

	defer wg.Done()
	count := 0

	for {

		mutex.Lock()
		if num <= 0 {
			fmt.Printf("%s：仓库空了\n", name)
			mutex.Unlock()
		} else {
			// 模拟消费产品的时间
			time.Sleep(time.Second * 3)
			count++
			num--

			fmt.Printf(
				"%s：消费了一件产品，我已经消费了%d件产品，当前仓库剩余容量为 %d\n",
				name,
				count,
				MaxNum-num,
			)
			mutex.Unlock()
		}

	}

}
