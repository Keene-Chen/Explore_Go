package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"
)

func MQTTSub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// fmt.Printf("Subscribed to topic: %s\n", topic)
	color.Cyan("Subscribed to topic: %s", topic)
}
