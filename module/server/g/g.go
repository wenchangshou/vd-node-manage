package g

import (
	"github.com/wenchangshou/vd-node-manage/common/Event"
	"runtime"
)

var (
	BinaryName string
	Version    string
	GitCommit  string
	GRedis     *Event.RedisClient
)

func VersionMsg() string {
	return Version + "@" + GitCommit
}
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
