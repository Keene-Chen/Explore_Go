package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	figure "github.com/common-nighthawk/go-figure"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	color "github.com/fatih/color"
)

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// fmt.Printf("\ntopic: %s\nrecmsg: %s\n", msg.Topic(), msg.Payload())
	color.Green("\ntopic: %s\nrecmsg: %s\n", msg.Topic(), msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client) {
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

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// fmt.Printf("Subscribed to topic: %s\n", topic)
	color.Cyan("Subscribed to topic: %s", topic)
}

func listenForKeyPress(client mqtt.Client) {
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

func main() {
	var broker = "gcloud.keenechen.cn"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("keenechen")
	opts.SetPassword("keenechen123*")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.ConnectTimeout = 3 * time.Second

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	// 创建一个 Figure 对象，传入要显示的文本
	myFigure := figure.NewFigure("KEENECHEN", "slant", true)
	myFigure.Print()

	// 启动一个协程监听键盘输入
	go listenForKeyPress(client)

	// 保持程序运行以持续接收消息
	select {}
}
