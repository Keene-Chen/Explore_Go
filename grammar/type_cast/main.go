package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to float32
	var a int = 1
	fmt.Printf("%f\n", float32(a))
	fmt.Printf("%f\n", float64(a))

	// float32 to int 丢失精度
	var b float32 = 1.1
	fmt.Printf("%d\n", int(b))

	// int to string
	var c int = 2
	fmt.Println(strconv.Itoa(c))

	// string to int
	var d string = "10"
	fmt.Println(strconv.Atoi(d))
}
