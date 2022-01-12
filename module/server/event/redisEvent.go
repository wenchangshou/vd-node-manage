package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/wenchangshou/vd-node-manage/common/Event"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"sync"
	"time"
)

type RedisEvent struct {
	client                *Event.RedisClient
	messageReplyIdChannel *sync.Map
}

func (manage RedisEvent) PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error) {
	var (
		n   int
		err error
	)
	b, _ := json.Marshal(msg)
	if n, err = manage.client.Publish(topic, b); err != nil {
		return "", fmt.Errorf("redis publish msg error:%s", err.Error())
	}
	if n <= 0 {
		return "", errors.New("设备未在线")
	}
	if reply {
		c := make(chan string)
		manage.messageReplyIdChannel.Store(msg.EventID, c)
		d := time.After(5 * time.Second)
		for {
			select {
			case r := <-c:
				manage.messageReplyIdChannel.Delete(msg.EventID)
				return r, nil
			case <-ctx.Done():
				manage.messageReplyIdChannel.Delete(msg.EventID)
				return "", nil
			case <-d:
				return "", errors.New("接收redis超时")

			}
		}
	}
	return "", nil
}
func (manage RedisEvent) ReplyMessage(id string, msg string) error {
	manage.messageReplyIdChannel.Range(func(key, value interface{}) bool {
		return true
	})
	c, ok := manage.messageReplyIdChannel.Load(id)
	if !ok {
		return nil
	}
	go func() {
		c.(chan string) <- msg
	}()
	return nil
}
func (manage RedisEvent) ServerEventMonitor(channel string, message []byte) error {
	reply := model.EventReply{}
	err := json.Unmarshal(message, &reply)
	if err != nil {
		return err
	}
	return manage.ReplyMessage(reply.EventID, reply.Body)
}
func (manage RedisEvent) Run() {
	for {

		if err := manage.client.Subscribe(context.TODO(), manage.ServerEventMonitor, "server"); err != nil {
			fmt.Println("redis disable", err)
		}

		time.Sleep(1 * time.Minute)
	}
}

func NewRedisEvent(args map[string]interface{}) (*RedisEvent, error) {
	var (
		addr   string
		db     int
		passwd string
		exists bool
	)
	if addr, exists = args["address"].(string); !exists {
		return nil, errors.New("address not exists")
	}
	passwd, _ = args["passwd"].(string)
	db = cast.ToInt(args["passwd"])
	c := Event.NewRedisClient(addr, db, passwd)
	e := RedisEvent{
		client:                c,
		messageReplyIdChannel: &sync.Map{},
	}
	go e.Run()
	return &e, nil
}
