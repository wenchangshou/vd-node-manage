package g

import "runtime"

var (
	BinaryName string
	Version    string
	GitCommit  string
)

func VersionMsg() string {
	return Version + "@" + GitCommit
}
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}
