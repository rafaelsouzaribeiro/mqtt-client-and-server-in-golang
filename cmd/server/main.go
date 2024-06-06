package main

import (
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/pkg/mqtt/server"
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/pkg/payload"
)

func main() {
	svc := server.NewBroker("localhost", 1883)

	svc.SetServer(&payload.Payload{
		Username: "root",
		Password: "123mudar",
		Topic:    "topic/test",
		Message:  "Test Message",
	})

}
