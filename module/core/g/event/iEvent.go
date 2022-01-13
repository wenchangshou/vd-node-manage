package event

import (
	"context"
	"errors"
)

var (
	GEvent IEvent
)

type ConsumeFunc func(channel string, message []byte) error

type IEvent interface {
	Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error
}

func NewEvent(provider string, maps map[string]interface{}) error {
	var (
		err error
	)
	if provider == "redis" {
		GEvent, err = NewRedisEvent(maps)
		if err != nil {
			return err
		}
		return nil
	} else if provider == "nats" {

	} else {
		return errors.New("unknown event provider")
	}
	return nil
}
