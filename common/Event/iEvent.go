package Event

import "context"

type Driver interface {
	Publish(channel, message string) (int, error)
	Subscribe(ctx context.Context, consumeFunc ConsumeFunc, channel ...string) error
}
