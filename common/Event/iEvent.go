package Event

import (
	"context"
	"errors"
	"github.com/wenchangshou/vd-node-manage/common/model"
)

type Driver interface {
	PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error)
	Subscribe(ctx context.Context, consumeFunc ConsumeFunc, channel ...string) error
}

func NewEvent(provider string, maps map[string]interface{}) (Driver, error) {
	if provider == "redis" {
		var (
			address  string
			db       float64
			password string
			exists   bool
		)
		if address, exists = maps["address"].(string); !exists {
			return nil, errors.New("address not exists")
		}
		if db, exists = maps["db"].(float64); !exists {
			return nil, errors.New("db not exists")
		}
		if password, exists = maps["password"].(string); !exists {
			return nil, errors.New("password not exists")
		}
		r := NewRedisClient(address, int(db), password)
		return r, nil
	} else if provider == "nats" {
		r, err := NewNatsEvent(maps)
		return r, err
	}
	return nil, errors.New("unknown event provider")
}
