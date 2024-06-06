package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/pkg/payload"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func (b *Broker) SetClient(pay payload.Payload, canalChan chan<- payload.Payload) {

	var broker = b.Broker
	var port = b.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetUsername(pay.Username)
	opts.SetPassword(pay.Password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Subscribe(pay.Topic, 1, func(c mqtt.Client, m mqtt.Message) {
		canalChan <- payload.Payload{
			Topic:     m.Topic(),
			Message:   string(m.Payload()),
			MessageId: m.MessageID(),
		}
	})

	b.Client = client
	b.Topic = pay.Topic

	token.Wait()
	fmt.Printf("Subscribed to topic: %s", pay.Topic)

}

func (b *Broker) PublishMessage(message string) {
	token := b.Client.Publish(b.Topic, 1, false, message)
	token.Wait()
	fmt.Printf("Published message to topic: %s\n", b.Topic)
}
