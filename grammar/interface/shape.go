package main

import "math"

// 定义一个求矩形面积的接口
type Shape interface {
	area() float64
}

// 定义一个矩形结构体
type Ractangle struct {
	width  float64
	height float64
}

// 实现求矩形面积方法
func (ract Ractangle) area() float64 {
	return ract.width * ract.height
}

type Circle struct {
	radius float64
}

// 实现求圆形面积方法
func (circ Circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}
