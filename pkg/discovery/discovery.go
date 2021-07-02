package discovery

import (
	"encoding/json"
	"net"
	"time"

	"github.com/wenchangshou2/vd-node-manage/pkg/conf"
)

const maxBufferSize = 1024

func GetServerInfo() (data []byte, err error) {
	tmp := map[string]interface{}{}
	tmp["Service"] = "RegisterInfo"
	tmp["ip"] = conf.ServerConfig.Ip
	tmp["port"] = conf.ServerConfig.Port
	return json.Marshal(tmp)
}
func discoveryServer(conn net.PacketConn) {
	buffer := make([]byte, maxBufferSize)
	for {
		_, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			time.Sleep(time.Millisecond)
			continue
		}
		data, err := GetServerInfo()
		if err == nil {
			conn.WriteTo(data, addr)
		}
	}
}
func InitDiscovery(ip string, port int) (err error) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	pc, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return
	}
	go discoveryServer(pc)
	return
}
