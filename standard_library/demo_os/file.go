// @file file.go
// @desc file
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.05-15:06:11

package main

import (
	"fmt"
	"os"
)

func TestFile() {
	// 创建文件并写入内容
	err := os.WriteFile("example.txt", []byte("Beautiful, World!"), 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created and written successfully")

	// 读取文件内容
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 打印文件内容
	fmt.Printf("File content: %s\n", string(data))

	// 删除文件
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}
