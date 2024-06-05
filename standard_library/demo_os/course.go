// @file course.go
// @desc course
// @author KeeneChen <keenechen@qq.com>
// @since  2024.06.05-15:02:23

package main

import (
	"fmt"
	"os"
)

func TestCourse() {
	gid := os.Getgid()
	fmt.Println("Group ID:", gid)
	uid := os.Getuid()
	fmt.Println("User ID:", uid)
	ppid := os.Getppid()
	fmt.Println("Parent PID:", ppid)
	pid := os.Getpid()
	fmt.Println("Current PID:", pid)
}
