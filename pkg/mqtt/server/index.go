package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/rafaelsouzaribeiro/mqtt-client-server-golang/pkg/payload"
)

func (b *Broker) SetServer(pay *payload.Payload) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	opts := mqtt.Options{
		InlineClient: true,
	}

	server := mqtt.New(&opts)

	_ = server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{
				{Username: auth.RString(pay.Username), Password: auth.RString(pay.Password), Allow: true},
			}}})

	listener := listeners.Config{
		Address: fmt.Sprintf("%s:%d", b.Broker, b.Port),
	}

	tcp := listeners.NewTCP(listener)
	err := server.AddListener(tcp)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		var i = 0
		for range time.Tick(time.Second * 1) {
			message := fmt.Sprintf("%s %d!\n", pay.Message, i)
			err := server.Publish(pay.Topic, []byte(message), false, 0)
			if err != nil {
				server.Log.Error("server.Publish", "error", err)
			}
			server.Log.Info("main.go direct message to " + pay.Topic)
			i++
		}
	}()

	sigReceived := <-sigs
	server.Log.Info("Received signal", "signal", sigReceived)
	server.Close()
}
