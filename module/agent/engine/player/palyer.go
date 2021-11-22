package player

import (
	"errors"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"sync"
)

type IPlayer interface {
	Open(*sync.WaitGroup, int) error
	GetThreadId() uint32
	Close() error
	// Check  error
	Check() (bool, error)
	// OpenCheck() (bool, error)
}

func MakePlayer(windowInfo e.Window, params string, service string, source string) (IPlayer, error) {
	// 先处理标准player
	playerPath := GetPlayerPath(service)
	if playerPath == "" {
		return nil, errors.New("未找到播放器")
	}
	if service == "ppt" || service == "video" || service == "pdf" || service == "image" || service == "http" {
		resourcePlayer := ResourcePlayer{
			Window:   windowInfo,
			PlayPath: playerPath,
			Source:   source,
		}
		return &resourcePlayer, nil
	}
	return nil, nil
}
