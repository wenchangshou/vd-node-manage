package g

import "sync"

// MessageChannel 消息中转通道
type MessageChannel struct {
	sync.Mutex
}
