package Event

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

type NatsEvent struct {
	nc *nats.Conn
}

func (client NatsEvent) fanOut(channel []string) chan []byte {
	c := make(chan []byte)
	return c
}
func (client NatsEvent) Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error {

	for _, v := range channel {
		go func(channel string) {
			client.nc.Subscribe(channel, func(msg *nats.Msg) {
				reply, err := consume(channel, msg.Data)
				if err != nil {
					client.nc.Publish(msg.Reply, reply)
				}
			})
		}(v)
	}
	return nil
}
func NewNatsEvent(maps map[string]interface{}) (*NatsEvent, error) {
	var (
		url string
		// name string
		exists bool
	)
	if url, exists = maps["address"].(string); !exists {
		return nil, errors.New("address not exists")
	}
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("connect nats fail:%s", err.Error())
	}
	c := NatsEvent{
		nc,
	}
	return &c, nil
}
