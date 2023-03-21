package mqtt

import (
	"door/config"
	"door/global"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

var Client mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	log.Println("Received message: " + payload)
	if payload == "pong" {
		global.DeviceLastOnlineTime = time.Now().Unix()
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func connect() mqtt.Client {
	log.Println("Connecting to " + config.Mqtt.Addr)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Mqtt.Addr)

	opts.SetClientID(config.Mqtt.ClientID)
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe
	client.Subscribe(config.Mqtt.Topic, 1, messagePubHandler)
	return client
}

func Init() {
	Client = connect()
	Publish(config.Mqtt.Topic, "ping")
}

func Publish(topic, data string) {
	if !Client.IsConnected() {
		Client = connect()
	}
	token := Client.Publish(topic, 0, false, data)
	token.Wait()
}

//func publish(client mqtt.Client) {
//	num := 10
//	for i := 0; i < num; i++ {
//		text := fmt.Sprintf("Message %d", i)
//		token := client.Publish("topic/test", 0, false, text)
//		token.Wait()
//		time.Sleep(time.Second)
//	}
//}

//func sub(client mqtt.Client) {
//	topic := "servo"
//	token := client.Subscribe(topic, 1, nil)
//	token.Wait()
//	fmt.Printf("Subscribed to topic: %s", topic)
//}
