package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"

	"github.com/wenchangshou/vd-node-manage/module/server/g"
)

type Device int
type Task int
type Event int
type Resource int

func Start() {
	addr := g.Config().Listen
	if g.Config().Mode == "docker" {
		addr = os.Getenv("LISTEN_ADDR") + ":6030"
	}
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
