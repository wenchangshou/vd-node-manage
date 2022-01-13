package g

import (
	"log"
	"net"
	"strings"
	"time"
)

var LocalIp string

// InitLocalIp 初始化本地ip
func InitLocalIp() {
	serverInfo := GetServerInfo()
	if serverInfo.Register {
		conn, err := net.DialTimeout("tcp", serverInfo.Http.Address, time.Second*10)
		if err != nil {
			log.Println("get local addr failed!")
		} else {
			LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
		}
	} else {
		log.Println("heartbeat is not enabled,can't get local addr")
	}
}
