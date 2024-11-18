package main

import (
	"fmt"
	"time"
)

/**
编写一个Goroutine，定时打印当前时间
*/

// var wg sync.WaitGroup

func init() {
	// wg.Add(1)
}

func main() {

	go printTimeTicker2(time.Second * 5)

	for {
		// 其他业务信息
	}

	// wg.Wait()
}

// printTimeTicker 按照指定时间间隔[duration]打印当前时间
func printTimeTicker1(duration time.Duration) {

	tick := time.Tick(duration)

	for t := range tick {
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	}

	// wg.Done()
}

func printTimeTicker2(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t.Format("2006-01-02 15:04:05"))
		}
	}
	// wg.Done()
}
