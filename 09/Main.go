package main

import (
	"fmt"
	myerror "go-exercises/09/error"
)

func main() {

	var num int

	fmt.Print("请输入一个整数： ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return
	}

	fmt.Println(num)

	if num <= 0 {
		myError := myerror.NewMyError(101, "num小于等于0")
		fmt.Println(myError)
	}

}
