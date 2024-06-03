package main

import (
	"github.com/rafaelsouzaribeiro/mqtt-client-server-golang/pkg/mqtt/server"
	"github.com/rafaelsouzaribeiro/mqtt-client-server-golang/pkg/payload"
)

func main() {
	svc := server.NewBroker("localhost", 1883)

	svc.SetServer(&payload.Payload{
		Username: "root",
		Password: "123mudar",
		Topic:    "topic/test",
		Message:  "Message Test",
	})

}
