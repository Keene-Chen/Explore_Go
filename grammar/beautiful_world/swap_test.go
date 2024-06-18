package main

import (
	"fmt"
	"testing"
)

func swap(x, y string) (string, string) {
	return y, x
}

func Test3(t *testing.T) {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}
