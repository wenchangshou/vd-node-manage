package g

import (
	"errors"
	"github.com/toolkits/net"
	"math"
	"net/rpc"
	"sync"
	"time"
)

type SingleConnRpcClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}
var (
	ServerRpcClient *SingleConnRpcClient
)
func InitRpcClients() {
	if Config().Server.Register{
		ServerRpcClient=&SingleConnRpcClient{
			rpcClient: nil,
			RpcServer: Config().Server.RpcAddress,
			Timeout:   time.Duration(Config().Heartbeat.Timeout)*time.Millisecond,
		}
	}
}
func (client *SingleConnRpcClient) close() {
	if client.rpcClient != nil {
		client.rpcClient.Close()
		client.rpcClient = nil
	}
}
func (client *SingleConnRpcClient) serverConn() error {
	if client.rpcClient != nil {
		return nil
	}
	var err error
	var retry int = 1
	for {
		if client.rpcClient != nil {
			return nil
		}
		client.rpcClient, err = net.JsonRpcClient("tcp", client.RpcServer, client.Timeout)
		if err != nil {
			if retry > 3 {
				return err
			}
			time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)
			retry++
			continue
		}
		return err
	}
}

func (client *SingleConnRpcClient) Call(method string, args interface{}, reply interface{}) error {
	client.Lock()
	defer client.Unlock()
	err := client.serverConn()
	if err != nil {
		return err
	}
	timeout := time.Duration(100 * time.Second)
	done := make(chan error, 1)
	go func() {
		err := client.rpcClient.Call(method, args, reply)
		done <- err
	}()
	select {
	case <-time.After(timeout):
		client.close()
		return errors.New(client.RpcServer + " rpc call timeout")
	case err := <-done:
		if err != nil {
			client.close()
			return err
		}

	}
	return nil
}
