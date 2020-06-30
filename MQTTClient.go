package ModbusMQTT

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Connect(clientId string, uri string) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	//opts.SetUsername()
	//opts.SetPassword()
	opts.SetClientID(clientId)
	return opts
}

func Listen(uri string, topic string) {
	client := Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("[%s]: %s\n", msg.Topic(), string(msg.Payload()))
		fmt.Println("dut")
	})
}

func MQTTPublish(client mqtt.Client,topic string,payload interface{}){
	client.Publish(topic,0,true,payload)
}

