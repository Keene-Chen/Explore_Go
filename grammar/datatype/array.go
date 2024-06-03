package main

import (
	"fmt"
)

func fun1() {
	var arr1 [5]int
	arr1 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1[%d]: %d ", i, arr1[i])
	}
	fmt.Println()
}

func fun2() {
	var arr2 = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d]: %d ", i, arr2[i])
	}
	fmt.Println()
}

func fun3() {
	arr3 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d]: %d ", i, arr3[i])
	}
	fmt.Println()
}

func fun4() {
	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr4[%d]: %d ", i, arr4[i])
	}
	fmt.Println()
}

func fun5() {
	arr5 := [5]int{1: 10, 3: 30}
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("arr5[%d]: %d ", i, arr5[i])
	}
	fmt.Println()
}

func fun6() {
	// float32 数组
	arr6 := [5]float32{1.1, 2.2, 3.3, 4.3, 5.5}
	for i := 0; i < len(arr6); i++ {
		fmt.Printf("arr6[%d]: %f ", i, arr6[i])
	}
	fmt.Println()
}

func fun7() {
	// string 数组
	arr7 := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr7); i++ {
		fmt.Printf("arr7[%d]: %s ", i, arr7[i])
	}
	fmt.Println()
}

func add(a []int, b []int) []int {
	var c []int
	for i := 0; i < len(a); i++ {
		c = append(c, a[i]+b[i])
	}
	return c
}

func main() {
	fun1()
	fun2()
	fun3()
	fun4()
	fun5()
	fun6()
	fun7()

	// 数组当做函数形参传递
	arr8 := add([]int{1, 2, 3}, []int{4, 5, 6, 7})
	for i := 0; i < len(arr8); i++ {
		fmt.Printf("arr8[%d]: %d ", i, arr8[i])
	}
	fmt.Println()

	a, b := 1, 1.1
	s :=
		fmt.Println(a + b)
}
