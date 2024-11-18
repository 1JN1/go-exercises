package main

import (
	"fmt"
	"go-exercises/02/car"
	"time"
)

func main() {

	t, err := time.Parse("2006-01-02", "2023-11-25")
	if err != nil {
		fmt.Println(err)
		return
	}

	param := make(map[string][]string)

	param["车长"] = []string{"15mm"}

	param["最高车速"] = []string{"15km/h"}

	param["测试"] = []string{"测试一", "测试二"}

	c := car.Car{
		Brand:          "宝马",
		Country:        "德国",
		Price:          3888888.88,
		Parameters:     param,
		ProductionDate: t,
	}

	fmt.Println(c.Info())

}
