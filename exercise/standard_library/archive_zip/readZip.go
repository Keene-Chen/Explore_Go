package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func readZip(zipPath string) {
	// 打开zip压缩文件
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatalf("Failed to open ZIP file: %v", err)
	}
	defer zipReader.Close()

	// 读取压缩文件
	for _, file := range zipReader.File {
		fmt.Printf("Reading file: %s\n", file.Name)

		// 打开压缩文件中的文件
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatalf("Failed to open file inside ZIP: %v", err)
		}

		// 创建目标文件路径
		targetPath := filepath.Join("./output", file.Name)

		// 如果是目录，则创建目录
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(targetPath, os.ModePerm)
			if err != nil {
				log.Fatalf("Failed to create directory: %v", err)
			}
			err = zippedFile.Close()
			if err != nil {
				return
			}
			continue
		}

		// 创建目标文件的目录
		err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}

		// 创建目标文件
		targetFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Fatalf("Failed to create target file: %v", err)
		}

		// 将压缩文件中的内容复制到目标文件
		_, err = io.Copy(targetFile, zippedFile)
		if err != nil {
			log.Fatalf("Failed to copy contents: %v", err)
		}

		// 立即关闭文件
		err = targetFile.Close()
		if err != nil {
			return
		}
		err = zippedFile.Close()
		if err != nil {
			return
		}
	}

	fmt.Println("ZIP extraction completed successfully.")
}
