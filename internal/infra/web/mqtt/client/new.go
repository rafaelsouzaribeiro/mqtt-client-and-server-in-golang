package client

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Broker struct {
	Broker string
	Port   int
	Client mqtt.Client
	Topic  string
}

func NewBroker(broker string, port int) *Broker {
	return &Broker{
		Broker: broker,
		Port:   port,
	}
}
