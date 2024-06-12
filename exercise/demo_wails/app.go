package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gek64/displayController"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	// 获取系统显示设备
	compositeMonitors, err := displayController.GetCompositeMonitors()
	if err != nil {
		log.Printf("Error getting composite monitors: %v", err)
		return fmt.Sprintf("Hello %s, unable to retrieve monitor information.", name)
	}

	var currentValue, min, max int

	for i, compositeMonitor := range compositeMonitors {
		fmt.Printf("Monitor No.%d\n", i)
		fmt.Printf("PhysicalInfo:%+v\n", compositeMonitor.PhysicalInfo)
		fmt.Printf("SysInfo:%+v\n", compositeMonitor.SysInfo)

	}

	currentValue, min, max, err = displayController.GetMonitorBrightness(compositeMonitors[0].PhysicalInfo.Handle)
	if err != nil {
		fmt.Printf("Error getting brightness for monitor: %v\n", err)
	}

	return fmt.Sprintf("Hello %s, Brightness: currentValue=%d, min=%d, max=%d", name, currentValue, min, max)
}

func (a *App) SetBrightness(value string) string {
	// 获取系统显示设备
	compositeMonitors, err := displayController.GetCompositeMonitors()
	if err != nil {
		log.Printf("Error getting composite monitors: %v", err)
	}
	num, _ := strconv.Atoi(value)
	err = displayController.SetMonitorBrightness(compositeMonitors[1].PhysicalInfo.Handle, num)
	if err != nil {
		log.Printf("Error setting brightness: %v", err)
	}
	return fmt.Sprintf("Brightness set to %s", value)
}
