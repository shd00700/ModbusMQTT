package ModbusMQTT

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId string,url string) mqtt.Client {
	opts := createClientOptions(clientId,url)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string,url string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(url)
	return opts
}

/*func listen(topic string) {
	client := connect("sub","tcp://broker.hivemq.com:1883")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}*/

func MQPublish(topic string,qos byte,retained bool){
	client := connect("pub","tcp://broker.hivemq.com:1883")
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, qos, retained, t.String())
	}
}
