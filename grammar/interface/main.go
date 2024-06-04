package main

import "fmt"

// 定义接口
type Phone interface {
	call()
}

// 定义IPhone结构体
type IPhone struct {
	id  int
	num int
}

// 实现IPhone接口的call方法
func (iphone IPhone) call() {
	fmt.Println("iphone interface")
}

// 定义XMPhone结构体
type XMPhone struct {
	id  int
	num int
}

// 实现xmphone接口的call方法
func (xmphone XMPhone) call() {
	fmt.Println("xmphone interface")
}

func main() {
	a := IPhone{id: 1, num: 123}
	a.call()
	b := XMPhone{id: 2, num: 321}
	b.call()

	var phone Phone

	phone = new(IPhone)
	phone.call()

	phone = new(XMPhone)
	phone.call()

	ract := Ractangle{width: 300, height: 400}
	fmt.Println(ract.area())

	circ := Circle{radius: 2}
	fmt.Println(circ.area())
}
