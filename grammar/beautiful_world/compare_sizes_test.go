package main

import (
	"fmt"
	"testing"
)

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Test2(t *testing.T) {
	var num1 int = 10
	var num2 int = 20
	var result = max(num1, num2)
	fmt.Println("Max value is: ", result)
}
