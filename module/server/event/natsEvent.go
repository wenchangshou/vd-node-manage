package event

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
	Address string `json:"address"`
	nc      *nats.Conn
}

func (event NatsEvent) PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error) {
	if event.nc.IsClosed() {
		return "", errors.New("nats is closed")
	}
	b, _ := json.Marshal(msg)
	if reply {
		msg, err := event.nc.Request(topic, b, time.Second)
		if err != nil {
			return "", err
		}
		return string(msg.Data), nil
	}
	return "", event.nc.Publish(topic, b)
}

func (event NatsEvent) Connect() error {
	nc, err := nats.Connect(event.Address)
	if err != nil {
		return err
	}
	event.nc = nc
	return nil
}

func NewNatsEvent(args map[string]interface{}) (*NatsEvent, error) {
	var (
		err error
	)
	address, exists := args["address"]
	if !exists {
		return nil, errors.New("address not exists")
	}
	event := NatsEvent{Address: address.(string)}
	if err = event.Connect(); err != nil {
		return nil, fmt.Errorf("connect nats fail:%s", err.Error())
	}
	return &event, nil
}
