package Event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"time"
)

type NatsEvent struct {
	nc *nats.Conn
}

func (client NatsEvent) PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error) {
	if client.nc.IsClosed() {
		return "", errors.New("nats is closed")
	}
	b, _ := json.Marshal(msg)
	if reply {
		msg, err := client.nc.Request(topic, b, time.Second)
		if err != nil {
			return "", err
		}
		return string(msg.Data), nil
	}
	return "", client.nc.Publish(topic, b)
}

//func (client NatsEvent) publish(channel string, message []byte) (int, error) {
//	if client.nc.IsClosed() {
//		return "", errors.New("nats is closed")
//	}
//	b, _ := json.Marshal(msg)
//	if reply {
//		msg, err := event.nc.Request(topic, b, time.Second)
//		if err != nil {
//			return "", err
//		}
//		return string(msg.Data), nil
//	}
//	return "", event.nc.publish(topic, b)
//}

func (client NatsEvent) fanOut(channel []string) chan []byte {
	c := make(chan []byte)
	return c
}

// Subscribe 订阅事件
func (client NatsEvent) Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error {
	for _, v := range channel {
		go func(channel string) {
			client.nc.Subscribe(channel, func(msg *nats.Msg) {
				reply, err := consume(channel, msg.Data)
				if err == nil && reply != nil {
					b := reply.Body
					client.nc.Publish(msg.Reply, []byte(b))
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
