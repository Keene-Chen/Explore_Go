package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"syscall"
)

func ListenForKeyPress(client mqtt.Client) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigChan:
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				client.Disconnect(250)
				fmt.Println("\nDisconnected from MQTT broker. Exiting program.")
				os.Exit(0)
			}
		default:
			var input string
			fmt.Scanln(&input)
			if input == "q" {
				client.Disconnect(250)
				fmt.Println("Program exited by user.")
				os.Exit(0)
			}
		}
	}
}
