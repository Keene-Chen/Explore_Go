package main

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"
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

	MQTTSub(client)
	MQTTPublish(client)

	// 创建一个 Figure 对象，传入要显示的文本
	myFigure := figure.NewFigure("KEENECHEN", "slant", true)
	myFigure.Print()

	// 启动一个协程监听键盘输入
	go ListenForKeyPress(client)

	// 保持程序运行以持续接收消息
	select {}
}
