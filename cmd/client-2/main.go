package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/mqtt-client-server-golang/pkg/mqtt/client"
	"github.com/rafaelsouzaribeiro/mqtt-client-server-golang/pkg/payload"
)

func main() {

	svc := client.NewBroker("localhost", 1883)

	go func() {
		channel := make(chan payload.Payload)
		svc.SetClient(payload.Payload{
			Username: "root",
			Password: "123mudar",
			Topic:    "topic/test",
		}, channel)

		for messages := range channel {
			fmt.Printf("Message: %s Topic: %s Message ID: %d \n",
				messages.Message, messages.Topic, messages.MessageId)

		}
	}()

	select {}
}
