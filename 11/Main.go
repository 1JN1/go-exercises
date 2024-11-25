package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)

	defer cancelFunc()

	go worker(ctx, 10086)

	time.Sleep(6 * time.Second)

	fmt.Println("主线程结束了")

}

func worker(ctx context.Context, id int) {

	for {

		select {
		case <-ctx.Done():
			fmt.Printf("worker %d 工作结束了！\n", id)
			return
		default:
			fmt.Printf("worker %d 正在工作中\n", id)
			time.Sleep(500 * time.Millisecond)
		}

	}

}
