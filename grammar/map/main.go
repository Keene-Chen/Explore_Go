package main

import "fmt"

func main() {
	// 定义map
	m := make(map[string]int)
	fmt.Println(m)

	// 初始化map
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m1)

	// 获取map中的值
	k, v := m1["a"]
	fmt.Println(k, v)

	// 获取map中所有键
	var i int
	keys := make([]string, len(m1))
	for k := range m1 {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// 修改值和删除值
	m1["c"] = 10
	delete(m1, "b")
	fmt.Println(m1)
}
