package discover

//const maxBufferSize = 1024
//
//var (
//	isRegister bool
//)
//
//// 注册返回的结构体
//type registerStr struct {
//	Service string `json:"service"`
//	Ip      string `json:"ip"`
//	Port    uint   `json:"port"`
//}
//
//// Invention 发现发现
//type Invention struct {
//	conn          *net.UDPConn
//	pc            *ipv4.PacketConn
//	address       *net.UDPAddr
//	regMsg        chan registerStr
//	regRequestMsg chan net.Addr
//	Ip            string
//	Port          uint
//}
//
//func NewInvention(ip net.IP, port uint) (*Invention, error) {
//	address := &net.UDPAddr{
//		IP:   ip,
//		Port: int(port),
//	}
//	fmt.Println("new addr", port)
//	conn, err := net.ListenUDP("udp4", address)
//	if err != nil {
//		return nil, err
//	}
//	pc := ipv4.NewPacketConn(conn)
//	return &Invention{
//		pc:            pc,
//		address:       address,
//		conn:          conn,
//		regMsg:        make(chan registerStr),
//		regRequestMsg: make(chan net.Addr, 5),
//	}, nil
//
//}
//func (invention Invention) MakePacket() []byte {
//	return []byte("hello")
//}
//func (invention Invention) IsRegisterMessage(msg []byte) bool {
//	m := invention.MakePacket()
//	for k, v := range msg {
//		if m[k] != v {
//			return false
//		}
//	}
//	return true
//}
//func (invention Invention) doBroadcast() error {
//	dstAddr := &net.UDPAddr{IP: net.IPv4bcast, Port: 8889}
//	b := invention.MakePacket()
//	fmt.Println("发送")
//	_, err := invention.conn.WriteTo(b, dstAddr)
//	return err
//}
//func (invention Invention) receiveBroadcasts(conn *net.UDPConn) {
//	buf := make([]byte, 65536)
//	for {
//		n, addr, err := conn.ReadFrom(buf)
//		fmt.Printf("接收222:%s\n", string(buf[:n]))
//		if err != nil {
//			return
//		}
//		content := buf[:n]
//		if invention.IsRegisterMessage(content) {
//			invention.regRequestMsg <- addr
//			continue
//		}
//		fmt.Println("2222")
//		var packet registerStr
//		err = json.Unmarshal(buf[:n], &packet)
//		if err != nil {
//			continue
//		}
//		invention.regMsg <- packet
//
//		fmt.Printf("接收到数据:字节(%d),数据:%v", n, string(buf[:n]))
//	}
//}
//func (invention Invention) MakeServerInfo(ip string, port uint) []byte {
//	s := registerStr{
//		Ip:      ip,
//		Port:    port,
//		Service: "RegisterInfo",
//	}
//	data, _ := json.Marshal(s)
//	return data
//}
//func (invention Invention) Server(serverInfo e.serverinfo) {
//	go invention.receiveBroadcasts(invention.conn)
//	for addr := range invention.regRequestMsg {
//		buf := invention.MakeServerInfo(serverInfo.Ip, serverInfo.Port)
//		fmt.Printf("发送11,buf:%s,addr:%s\n", buf, addr.String())
//		invention.conn.WriteTo(buf, addr)
//	}
//}
//
//// WaitRegister 等待注册
//func (invention Invention) WaitRegister() (*e.ServerInfo, error) {
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//	go invention.receiveBroadcasts(invention.conn)
//	for {
//		select {
//		case <-ticker.C:
//			if err := invention.doBroadcast(); err != nil {
//				return nil, err
//			}
//		case msg := <-invention.regMsg:
//			return &e.ServerInfo{Ip: msg.Ip, Port: msg.Port}, nil
//		}
//
//	}
//
//}
