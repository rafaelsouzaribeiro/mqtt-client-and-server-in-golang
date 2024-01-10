package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func main() {
	// Create signals channel to run server until interrupted
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	opts := mqtt.Options{
		InlineClient: true,
	}
	// Create the new MQTT Server.
	server := mqtt.New(&opts)

	// Allow all connections.
	_ = server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{ // Auth disallows all by default
				{Username: "root", Password: "123mudar", Allow: true},
			}}})

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP("t1", "localhost:1883", nil)
	err := server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		var i = 0
		for range time.Tick(time.Second * 1) {
			message := fmt.Sprintf("Hello World %d!\n", i)
			err := server.Publish("topic/test", []byte(message), false, 0)
			if err != nil {
				server.Log.Error("server.Publish", "error", err)
			}
			server.Log.Info("main.go direct message to topic/test")
			i++
		}
	}()

	// Run server until interrupted
	<-done

	// Cleanup
}
