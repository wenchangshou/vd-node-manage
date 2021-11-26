package g

import (
	"log"
	"net"
	"strings"
	"time"
)

var LocalIp string

func InitLocalIp() {
	if Config().Server.Register {
		conn, err := net.DialTimeout("tcp", Config().Heartbeat.Addr, time.Second*10)
		if err != nil {
			log.Println("get local addr failed!")
		} else {
			LocalIp = strings.Split(conn.LocalAddr().String(), ":")[0]
		}
	} else {
		log.Println("heartbeat is not enabled,can't get local addr")
	}
}
