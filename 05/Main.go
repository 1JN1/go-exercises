package main

import "fmt"

func main() {

	str := "aaaAAA@aa%5sad"

	m := countTheNumberOfEachChar(str)

	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}

}

// countTheNumberOfEachChar 统计字符串 str 中每个字符出现的次数
func countTheNumberOfEachChar(str string) map[string]int {

	result := make(map[string]int)

	for _, s := range str {
		result[string(s)]++
	}

	return result
}
