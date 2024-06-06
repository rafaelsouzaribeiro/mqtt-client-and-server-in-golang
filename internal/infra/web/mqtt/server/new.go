package server

type Broker struct {
	Broker string
	Port   int
}

func NewBroker(broker string, port int) *Broker {
	return &Broker{
		Broker: broker,
		Port:   port,
	}
}
