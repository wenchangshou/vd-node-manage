package rpc

import (
	"github.com/wenchangshou/vd-node-manage/module/server/g"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type Device int
type Task int
type Event int
type Resource int

func Start() {
	addr := g.Config().Listen
	server := rpc.NewServer()
	server.Register(new(Device))
	server.Register(new(Task))
	server.Register(new(Event))
	server.Register(new(Resource))
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalln("listen error:", e)
	} else {
		log.Println("listening", addr)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("listener accept fail:", err)
			time.Sleep(time.Duration(100) * time.Millisecond)
			continue
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
