package event

import "context"

var (
	GEvent IEvent
)

type IEvent interface {
	Subscribe(ctx context.Context)
}

func NewEvent(provider string, maps map[string]interface{}) {

}
