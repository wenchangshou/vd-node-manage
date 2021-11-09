package invention

import (
	"encoding/json"
	"fmt"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/e"
	"github.com/wenchangshou2/vd-node-manage/module/agent/pkg/logging"
	"golang.org/x/net/ipv4"

	"net"
	"time"

)

const maxBufferSize = 1024

var (
	isRegister bool
)

// 注册返回的结构体
type registerStr struct{
	Service string `json:"service"`
	Ip string `json:"ip"`
	Port int `json:"port"`
}
// Invention 发现发现
type Invention struct {
	conn *net.UDPConn

	pc *ipv4.PacketConn
	address *net.UDPAddr
	regMsg chan registerStr
}

func NewInvention(ip string, port int) (*Invention,error) {
	//conn:=&net.UDPAddr{IP: net.IPv4bcast,Port: 8889}
	address:=&net.UDPAddr{
		IP: net.IPv4zero,
		Port:8889,
	}
	conn, err := net.ListenUDP("udp4", address)
	if err!=nil{
		return nil,err
	}
	pc:=ipv4.NewPacketConn(conn)
	return &Invention{
		pc: pc,
		address: address,
		conn: conn,
		regMsg: make(chan registerStr),
	},nil

}
func (invention Invention)MakePacket()[]byte{
	return []byte("hello")
}
func (invention Invention)doBroadcast()error{
	dstAddr := &net.UDPAddr{IP: net.IPv4bcast, Port: 8889}
	b:= invention.MakePacket()
	_,err:=invention.conn.WriteTo(b,dstAddr)
	return  err
}
func (invention Invention)receiveBroadcasts(conn *net.UDPConn){
	buf:=make([]byte,65536)
	for  {
		n,err:=invention.conn.Read(buf)
		if err!=nil{
			return
		}
		var packet registerStr
		err=json.Unmarshal(buf[:n],&packet)
		if err!=nil{
			continue
		}
		invention.regMsg<-packet

		fmt.Printf("接收到数据:字节(%d),数据:%v",n,string(buf[:n]))
	}
}
// WaitRegister 等待注册
func (invention Invention)WaitRegister()(*e.ServerInfo,error){
	ticker:=time.NewTicker(time.Second)
	defer ticker.Stop()
	go  invention.receiveBroadcasts(invention.conn)
	for  {
		select {
		case <-ticker.C:
			logging.GLogger.Info("发送注册请求")
			if err:=invention.doBroadcast();err!=nil{
				return nil,err
			}
			case msg:=<-invention.regMsg:
				logging.GLogger.Info(fmt.Sprintf("接收到注册消息:IP:%s,Port%d\n",msg.Ip,msg.Port))
				return &e.ServerInfo{Ip:msg.Ip,Port: msg.Port},nil
		}

	}

}
//func InitInvention() (serverInfo e.ServerInfo, err error) {
//	var (
//		pc net.PacketConn
//	)
//	registerChannel := make(chan e.ServerInfo)
//	addr := net.UDPAddr{
//		Port: 22213,
//		IP:   net.IPv4zero,
//	}
//	if pc, err = net.ListenUDP("udp", &addr); err != nil {
//		return
//	}
//	go writePump(pc)
//	go upnpClient(registerChannel, pc)
//	serverInfo = <-registerChannel
//	return
//}
//
//// upnpClient 自动发现服务
//func upnpClient(registerChannel chan e.ServerInfo, conn net.PacketConn) {
//	buffer := make([]byte, maxBufferSize)
//	message := make(map[string]interface{})
//	for {
//		n, _, err := conn.ReadFrom(buffer)
//		if err != nil {
//			continue
//		}
//
//		if err = json.Unmarshal(buffer[:n], &message); err == nil {
//			fmt.Println("message", message)
//			Service := message["Service"].(string)
//			if strings.Compare(Service, "RegisterInfo") == 0 {
//				ip := message["ip"].(string)
//				port := message["port"].(float64)
//				isRegister = true
//				serverInfo := e.ServerInfo{
//					Ip:   ip,
//					Port: int(port),
//				}
//				registerChannel <- serverInfo
//				return
//			}
//		} else {
//			fmt.Println("err", err)
//		}
//	}
//}
//
//func writePump(conn net.PacketConn) {
//	dstAddr := &net.UDPAddr{IP: net.IPv4bcast, Port: 8889}
//	for {
//		if isRegister {
//			return
//		}
//		fmt.Println("write pump")
//		conn.WriteTo([]byte("Hello"), dstAddr)
//		time.Sleep(2 * time.Second)
//	}
//}
