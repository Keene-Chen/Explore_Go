package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MQTTPublish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		// text := fmt.Sprintf("testMsg %d", i)
		// token := client.Publish("topic/test", 0, false, text)
		// token.Wait()
		// time.Sleep(time.Second)

		// 创建要发送的消息
		msg := Message{
			ID:      i,
			Content: fmt.Sprintf("testMsg %d", i),
		}

		// 序列化消息为 JSON
		jsonData, err := json.Marshal(msg)
		if err != nil {
			fmt.Printf("Error serializing message: %v\n", err)
			continue
		}

		// 发布 JSON 消息
		token := client.Publish("topic/test", 0, false, jsonData)
		token.Wait()
	}
}
