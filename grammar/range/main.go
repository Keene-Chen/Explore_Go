package main

import "fmt"

func main() {
	// 遍历切片
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Print(i, v)
	}
	fmt.Println()

	// 遍历数组
	arr := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Print(i, v)
	}

	// 遍历字符串
	str := "beautiful world"
	for i, v := range str {
		fmt.Print(i, string(v))
	}
	fmt.Println()
	for _, v := range str {
		fmt.Print(string(v))
	}
	fmt.Println()

	// 遍历map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Print(k, v)
	}
	fmt.Println()
}
