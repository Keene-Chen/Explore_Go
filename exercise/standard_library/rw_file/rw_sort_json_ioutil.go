package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

/**
 * Deprecated: Use SortJSON() instead
 */
func SortJSON(readFile string, writeFile string) {
	// 读取JSON文件
	fdata, err := ioutil.ReadFile(readFile)
	if err != nil {
		log.Fatal(err)
	}

	// 反序列化JSON数据到Data结构体切片
	var data []Data
	err = json.Unmarshal(fdata, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// 排序前打印数据
	fmt.Println("排序前的数据:")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %+v\n", i, data[i])
	}

	// 对数据进行排序
	sort.Sort(DataSlice(data))

	// 返回排序后的新数组
	sortedData := data
	fmt.Println("\n排序后的新数组:")
	for i := 0; i < len(sortedData); i++ {
		fmt.Printf("sortedData[%d]: %+v\n", i, sortedData[i])
	}

	// 将排序后的数据写入到新的JSON文件
	sortedDataJSON, err := json.MarshalIndent(sortedData, "", "  ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
	}

	// 写入到新的JSON文件
	err = ioutil.WriteFile(writeFile, sortedDataJSON, 0644)
	if err != nil {
		log.Fatal("Error during WriteFile(): ", err)
	}
}
