package main

import "fmt"

func main() {
	// 定义一个切片
	var slice []int
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 使用make函数创建切片
	slice1 := make([]int, 5)
	fmt.Println(slice1)

	// 使用数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice2 := arr[1:4]
	fmt.Println(slice2)

	// 切片的长度和容量
	slice3 := make([]int, 5, 10)
	fmt.Println(len(slice3), cap(slice3))

	// 切片截取
	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice4[1:3], slice4[:3], slice4[1:])
}
