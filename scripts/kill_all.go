package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func KillAll1() {
	// 获取所有进程列表
	cmd := exec.Command("tasklist")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing tasklist:", err)
		return
	}

	// 解析并找到所有名为go.exe的进程
	processes := strings.Split(out.String(), "\n")
	for _, process := range processes {
		if strings.Contains(process, "go.exe") {
			fields := strings.Fields(process)
			if len(fields) > 1 {
				pid := fields[1]
				killCmd := exec.Command("taskkill", "/F", "/PID", pid)
				err := killCmd.Run()
				if err != nil {
					fmt.Println("Error killing process:", err)
				} else {
					fmt.Println("Killed process with PID:", pid)
				}
			}
		}
	}
}
