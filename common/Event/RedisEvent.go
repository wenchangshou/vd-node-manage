package Event

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type ConsumeFunc func(channel string, message []byte) error
type RedisClient struct {
	pool *redis.Pool
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
		pool: pool,
	}
	return client
}

func (r *RedisClient) Publish(channel, message string) (int, error) {
	c := r.pool.Get()
	defer c.Close()
	n, err := redis.Int(c.Do("PUBLISH", channel, message))
	if err != nil {
		return 0, fmt.Errorf("redis publish %s %s, %v", channel, message, err)
	}
	return n, nil
}

// Subscribe 阅读服务端事件
func (r *RedisClient) Subscribe(ctx context.Context, consume ConsumeFunc, channel ...string) error {
	psc := redis.PubSubConn{Conn: r.pool.Get()}
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
				if err := consume(msg.Channel, msg.Data); err != nil {
					done <- err
					return
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
	return nil
}
