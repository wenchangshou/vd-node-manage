package event

import (
	"context"
	"errors"

	event2 "github.com/wenchangshou/vd-node-manage/common/Event"
)

type RedisEvent struct {
	client *event2.RedisClient
}
func(client RedisEvent) 	Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error{
	return nil
}


func NewRedisEvent(maps map[string]interface{}) (*RedisEvent, error) {
	var (
		address  string
		db       int
		password string
		exists   bool
	)
	if address, exists = maps["address"].(string); !exists {
		return nil, errors.New("addres not exists")
	}
	if db, exists = maps["db"].(int); !exists {
		return nil, errors.New("db not exists")
	}
	if password, exists = maps["password"].(string); !exists {
		return nil, errors.New("password not exists")
	}
	redisClient := event2.NewRedisClient(address, db, password)
	r := RedisEvent{
		client: redisClient,
	}
	return &r, nil
}
