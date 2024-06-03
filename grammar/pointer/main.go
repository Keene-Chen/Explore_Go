package main

import "fmt"

func add(x *int) {
	*x += 1
	fmt.Printf("%d\n", *x)
}

func main() {
	// &获取 int 指针
	var a int = 10
	fmt.Printf("a=%p\n", &a)

	// 定义指针变量,*int 表示指向 int 类型的指针,*p 表示指针变量的值
	var p *int = &a
	fmt.Printf("p=%d\n", *p)

	// 定义一个空指针
	var p1 *int
	fmt.Printf("p1=%p\n", p1)
	if p1 == nil {
		fmt.Println("p1 is nil")
	}

	// 指针数组
	var arr = [3]int{1, 2, 3}
	var pArr [3]*int
	for i := 0; i < len(arr); i++ {
		pArr[i] = &arr[i]
	}
	for i := 0; i < len(pArr); i++ {
		fmt.Printf("%d\n", *pArr[i])
	}

	// 数组指针
	// var pArr1 *[3]int = &arr
	pArr1 := &arr
	fmt.Printf("%d\n", (*pArr1)[0])

	// 指针传递给函数
	var b int = 10
	add(&b)

	// 二级指针
	var c int = 10
	var p2 *int = &c
	var p3 **int = &p2
	fmt.Printf("%d\n", **p3)
}
