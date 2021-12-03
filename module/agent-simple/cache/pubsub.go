package cache

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"time"
)

func InitSubPub(ctx context.Context, redisServerAddr string,
	onStart func() error,
	onMessage func(channel string, data []byte) error,
	channels ...string,
) error {
	const healthCheckPeriod = time.Minute
	c, err := redis.Dial("tcp", redisServerAddr,
		redis.DialReadTimeout(healthCheckPeriod+10*time.Second),
		redis.DialWriteTimeout(10*time.Second))
	if err != nil {
		return err
	}
	defer c.Close()
	psc := redis.PubSubConn{Conn: c}
	if err := psc.Subscribe(channels); err != nil {
		return err
	}
	done := make(chan error, 1)
	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- n
				return
			case redis.Message:
				if err := onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				switch n.Count {
				case len(channels):
					if err := onStart(); err != nil {
						done <- err
						return
					}
				case 0:
					done <- nil
					return
				}
			}
		}
	}()
	ticker := time.NewTicker(healthCheckPeriod)
	defer ticker.Stop()
loop:
	for {

		select {
		case <-ticker.C:
			if err = psc.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			break loop
		case err := <-done:
			return err
		}
	}
	if err := psc.Unsubscribe(); err != nil {
		return err
	}
	return <-done
}
