// @file clean.go
// @desc 清理当前目录下的tmp目录
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.05-10:50:55

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}

	// 遍历当前目录下的所有文件和目录
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		// 如果找到名为 "tmp" 的目录，并且它是一个目录
		if info.IsDir() && info.Name() == "tmp" {
			// 删除该目录
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println("删除目录失败:", err)
				return err
			}
			fmt.Println("已删除目录:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录失败:", err)
	}
}
