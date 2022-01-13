package Event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wenchangshou/vd-node-manage/common/model"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type ConsumeFunc func(channel string, message []byte) (*model.EventReply, error)
type RedisClient struct {
	messageReplyIdChannel *sync.Map
	pool                  *redis.Pool
}

func (client *RedisClient) PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error) {
	var (
		n   int
		err error
	)
	b, _ := json.Marshal(msg)
	if n, err = client.publish(topic, b); err != nil {
		return "", fmt.Errorf("redis publish msg error:%s", err.Error())
	}
	if n <= 0 {
		return "", errors.New("设备未在线")
	}
	if reply {
		c := make(chan string)
		client.messageReplyIdChannel.Store(msg.EventID, c)
		d := time.After(5 * time.Second)
		for {
			select {
			case r := <-c:
				client.messageReplyIdChannel.Delete(msg.EventID)
				return r, nil
			case <-ctx.Done():
				client.messageReplyIdChannel.Delete(msg.EventID)
				return "", nil
			case <-d:
				client.messageReplyIdChannel.Delete(msg.EventID)
				return "", errors.New("接收redis超时")

			}
		}
	}
	return "", nil
}
func (client RedisClient) ReplyMessage(id string, msg string) error {
	client.messageReplyIdChannel.Range(func(key, value interface{}) bool {
		return true
	})
	c, ok := client.messageReplyIdChannel.Load(id)
	if !ok {
		return nil
	}
	go func() {
		c.(chan string) <- msg
	}()
	return nil
}
func (client RedisClient) ServerEventMonitor(channel string, message []byte) (*model.EventReply, error) {
	reply := model.EventReply{}
	err := json.Unmarshal(message, &reply)
	if err != nil {
		return nil, err
	}
	return nil, client.ReplyMessage(reply.EventID, reply.Body)
}
func (client *RedisClient) Run() {
	for {

		if err := client.Subscribe(context.TODO(), client.ServerEventMonitor, "server"); err != nil {
			fmt.Println("redis disable", err)
		}

		time.Sleep(1 * time.Minute)
	}
}
func NewRedisClient(addr string, db int, passwd string) *RedisClient {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialPassword(passwd), redis.DialDatabase(db))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	client := &RedisClient{
		messageReplyIdChannel: &sync.Map{},
		pool:                  pool,
	}
	go client.Run()
	return client
}

func (client *RedisClient) publish(channel string, msg []byte) (int, error) {
	c := client.pool.Get()
	defer c.Close()
	n, err := redis.Int(c.Do("PUBLISH", channel, msg))
	if err != nil {
		return 0, fmt.Errorf("redis publish %s %s, %v", channel, msg, err)
	}
	return n, nil
}

// Subscribe 阅读服务端事件
func (client *RedisClient) Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error {
	psc := redis.PubSubConn{Conn: client.pool.Get()}
	defer psc.Close()
	if err := psc.Subscribe(redis.Args{}.AddFlat(channel)...); err != nil {
		return err
	}
	done := make(chan error, 1)
	go func() {
		for {
			switch msg := psc.Receive().(type) {
			case error:
				done <- fmt.Errorf("redis pubsub receive err:%v", msg)
				return
			case redis.Message:
				var (
					reply *model.EventReply
					err   error
				)
				if reply, err = consume(msg.Channel, msg.Data); err != nil {
					done <- err
					return
				}
				if reply != nil {
					b, _ := json.Marshal(reply)
					client.publish("server", b)
				}
			case redis.Subscription:
				if msg.Count == 0 {
					done <- nil
					return
				}
			}
		}
	}()
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			if err := psc.Unsubscribe(); err != nil {
				return fmt.Errorf("redis pubsub unsubscribe err:%v", err)
			}
			return nil
		case err := <-done:
			return err
		case <-tick.C:
			if err := psc.Ping(""); err != nil {
				return err
			}
		}
	}
}
