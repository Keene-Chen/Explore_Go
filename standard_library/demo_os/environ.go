// @file environ.go
// @desc environ
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.05-14:55:21

package main

import (
	"fmt"
	"os"
	"strings"
)

// GetEnviron gets all environment variables.
func GetEnviron() {
	envVars := os.Environ()
	for _, envVar := range envVars {
		kv := strings.SplitN(envVar, "=", 2)
		if len(kv) == 2 {
			fmt.Printf("%s=%s\n", kv[0], kv[1])
		} else {
			fmt.Printf("%s\n", kv[0])
		}
	}
}

// GetEnv gets the value of the specified environment variable.
// @param pathName string
// @return void
func GetEnv(pathName string) {
	path := os.Getenv(pathName)
	if path == "" {
		fmt.Println("No value found for key PATH")
		return
	}

	pathEntries := strings.Split(path, ";")

	for _, entry := range pathEntries {
		fmt.Println(entry)
	}
}

// SetEnv sets the value of the specified environment variable.
// @param key string
// @param value string
// @return void
func SetEnv(key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully set", key, "to", value)
}

func TestEnv() {
	hostName, _ := os.Hostname()
	fmt.Printf("主机名: %s\n", hostName)
	fmt.Println(os.Getpagesize())
}
