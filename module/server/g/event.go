package g

import "github.com/wenchangshou/vd-node-manage/common/Event"

var (
	GEvent Event.Driver
)

func InitEvent(provider string, args map[string]interface{}) error {
	var (
		err error
	)
	GEvent, err = Event.NewEvent(provider, args)
	return err
}
