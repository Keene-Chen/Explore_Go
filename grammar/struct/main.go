package main

import "fmt"

// 定义一个 Person 结构体
type Person struct {
	id    int
	name  string
	age   int
	score float32
}

func printPerson(p Person) {
	fmt.Printf("id=%d, name=%s, age=%d, score=%.2f\n", p.id, p.name, p.age, p.score)
}

func printPerson1(p *Person) {
	fmt.Printf("id=%d, name=%s, age=%d, score=%.2f\n", p.id, p.name, p.age, p.score)
}

func main() {
	stu := Person{101, "Tom", 20, 99.5}
	fmt.Printf("id=%d\n", stu.id)

	// 指定字段初始化
	stu1 := Person{id: 102, name: "Jerry", score: 98.5}
	fmt.Printf("name=%s\n", stu1.name)

	// 结构体作为函数参数
	printPerson(stu)

	// 结构体指针
	var p *Person = &stu
	fmt.Printf("id=%d\n", p.id)

	// 结构体指针作为函数参数
	printPerson1(p)

	// 结构体二级指针
	var q **Person = &p
	fmt.Printf("score=%.2f\n", (*q).score)

	// 结构体数组
	var arr [3]Person = [3]Person{{101, "Tom", 20, 99.5}, {102, "Jerry", 21, 98.5}, {103, "Lucy", 22, 97.5}}
	for i := 0; i < len(arr); i++ {
		printPerson(arr[i])
	}
}
