package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func main() {

	calculateTimeOfFunc(test01)

	// nums := []int{2, 1, 1, 2, 3, 1, 2, 5, 6, 7, 89, 1, 1, 6, 1, 4, 1, 235, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 2, 3, 1, 2, 5, 6, 7, 89, 1, 1, 6, 1, 4, 1, 235, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 2, 3, 1, 2, 5, 6, 7, 89, 1, 1, 6, 1, 4, 1, 235, 1, 1, 1, 1, 1, 2, 1}

	// calculateTimeOfFunc(test02[int], &nums, 1)

}

// calculateTimeOfFunc 计算函数f的运行时间
func calculateTimeOfFunc(f any, params ...any) {

	fnValue := reflect.ValueOf(f)

	if fnValue.Kind() != reflect.Func {
		panic("传入的不是一个函数")
	}

	// 将函数参数转换为反射类型的切片
	args := make([]reflect.Value, len(params))

	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}

	// 获取函数名
	name := runtimeFuncName(f)

	// 函数开始时间
	beginTime := time.Now()

	fnValue.Call(args)

	// 函数结束时间
	endTime := time.Now()

	fmt.Printf("函数%s{%s}的运行时间为%dms", name, fnValue.String(), endTime.Sub(beginTime).Milliseconds())
}

// runtimeFuncName 获取运行时函数名
func runtimeFuncName(f any) string {
	pc := reflect.ValueOf(f).Pointer()
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown function"
	}
	return fn.Name()
}

func test01() {

	for i := range 10000000 {
		fmt.Println(i)
	}

}

func test02[T int | float32 | float64 | string](nums *[]T, elem T) {

	for i := 0; i < len(*nums); i++ {

		if (*nums)[i] == elem {

			*nums = append((*nums)[:i], (*nums)[i+1:]...)
			i--
		}

	}
}
