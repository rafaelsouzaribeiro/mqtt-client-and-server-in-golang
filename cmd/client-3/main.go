package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/internal/infra/web/mqtt/client"
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/internal/usecase/dto"
)

func main() {
	PublishMessage(true)
}

func PublishMessage(clientMessage bool) {

	svc := client.NewBroker("localhost", 1883)

	go func() {
		channel := make(chan dto.Payload)
		svc.SetClient(dto.Payload{
			Username: "root",
			Password: "123mudar",
			Topic:    "topic/test",
		}, channel)

		if clientMessage {
			svc.PublishMessage("Client Message 3")
		}

		for messages := range channel {
			fmt.Printf("Message: %s Topic: %s Message ID: %d \n",
				messages.Message, messages.Topic, messages.MessageId)

		}

	}()

	select {}
}
