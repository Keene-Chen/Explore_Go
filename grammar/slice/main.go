package main

import "fmt"

func main() {
	// 定义一个切片 
	var slice []int
	fmt.Println(slice)

	// 使用make函数创建切片
	slice1 := make([]int, 5)
	fmt.Println(slice1)
}