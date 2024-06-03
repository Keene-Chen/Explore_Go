package main

// 定义一个Data结构体
type Data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// 定义一个类型DataSlice，用于实现sort.Interface
type DataSlice []Data

// 实现sort.Interface接口的Len方法
func (d DataSlice) Len() int {
	return len(d)
}

// 实现sort.Interface接口的Less方法
func (d DataSlice) Less(i, j int) bool {
	return d[i].Age < d[j].Age
}

// 实现sort.Interface接口的Swap方法
func (d DataSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	readFile := "./res/data.json"
	writeFile := "./res/sorted_data.json"
	writeFile1 := "./res/sorted_data1.json"

	SortJSON(readFile, writeFile)
	SortJSON1(readFile, writeFile1)
}
