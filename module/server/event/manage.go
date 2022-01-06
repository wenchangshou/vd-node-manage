package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/wenchangshou/vd-node-manage/common/Event"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	"sync"
	"time"
)

type Manage struct {
	client                *Event.RedisClient
	messageReplyIdChannel *sync.Map
}

func (manage Manage) Run() {
	for {

		if err := manage.client.Subscribe(context.TODO(), manage.ServerEventMonitor, "server"); err != nil {
			fmt.Println("redis disable", err)
		}

		time.Sleep(1 * time.Minute)
	}
}
func (manage *Manage) ReplyMessage(id string, msg string) error {
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

func (manage *Manage) ServerEventMonitor(channel string, message []byte) error {
	reply := model.EventReply{}
	err := json.Unmarshal(message, &reply)
	if err != nil {
		return err
	}
	return manage.ReplyMessage(reply.EventID, reply.Body)
}

func (manage *Manage) PublishEvent(ctx context.Context, action string, topic string, body []byte, reply bool) (string, error) {
	var (
		msg model.EventRequest
		n   int
	)
	msg.Action = action
	msg.Arguments = body
	msg.Reply = reply
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	msg.EventID = uid.String()
	m, _ := json.Marshal(msg)
	if n, err = manage.client.Publish(topic, string(m)); err != nil {
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

var (
	GManage *Manage
)

func InitEvent(cfg *g.CacheConfig) {
	c := Event.NewRedisClient(cfg.Addr, cfg.DB, cfg.Passwd)
	GManage = &Manage{
		client:                c,
		messageReplyIdChannel: &sync.Map{},
	}
	go GManage.Run()
}
