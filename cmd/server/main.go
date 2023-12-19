package main

import (
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
	_ = server.AddHook(new(auth.AllowHook), nil)

	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP("t1", "0.0.0.0:1883", nil)
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
		for range time.Tick(time.Second * 1) {
			err := server.Publish("topic/test", []byte("Hello World!\n"), false, 0)
			if err != nil {
				server.Log.Error("server.Publish", "error", err)
			}
			server.Log.Info("main.go issued direct message to direct/publish")
		}
	}()

	// Run server until interrupted
	<-done

	// Cleanup
}
