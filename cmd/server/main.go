package main

import (
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/internal/infra/web/mqtt/server"
	"github.com/rafaelsouzaribeiro/mqtt-client-and-server-in-golang/internal/usecase/dto"
)

func main() {
	svc := server.NewBroker("localhost", 1883)

	svc.SetServer(&dto.Payload{
		Username: "root",
		Password: "123mudar",
		Topic:    "topic/test",
		Message:  "Test Message",
	})

}
