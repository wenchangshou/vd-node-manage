package event

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/wenchangshou/vd-node-manage/common/model"
)

type IEvent interface {
	PublishEvent(ctx context.Context, topic string, msg *model.EventRequest, reply bool) (string, error)
}

var (
	GEvent IEvent
)

func GetEventCmd(action string, did uint, body []byte, reply bool) (*model.EventRequest, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	msg := model.EventRequest{
		EventID:   id.String(),
		Action:    action,
		DeviceID:  did,
		Arguments: body,
		Reply:     reply,
	}
	return &msg, nil

}
func NewEvent(provide string, args map[string]interface{}) error {
	var (
		err error
	)
	if provide == "nats" {
		if GEvent, err = NewNatsEvent(args); err != nil {
			return err
		}
	} else if provide == "redis" {
		if GEvent, err = NewRedisEvent(args); err != nil {
			return err
		}
	} else {
		return errors.New("unknown event provider:" + provide)
	}
	return nil
}
