package playerService

import (
	"context"
	"errors"
	"fmt"
	playerRpc "github.com/wenchangshou/vd-node-manage/module/core/engine/player/rpc"
	"google.golang.org/grpc"
	"time"
)

type RpcPlayerService struct {
	port int `json:"port"`
}

func (r RpcPlayerService) Ping() (bool, error) {

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", r.port), grpc.WithBlock())
	if err != nil {
		return false, errors.New("connect grpc server fail:" + err.Error())
	}
	defer conn.Close()
	c := playerRpc.NewRpcCallClient(conn)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	res, err := c.Ping(ctx, &playerRpc.EmptyMessage{})
	if err != nil {
		return false, errors.New("call remote rpc server fail:" + err.Error())
	}
	return res.Code == 0, nil
}

func (r RpcPlayerService) Control(payload string) (string, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", r.port), grpc.WithInsecure())
	if err != nil {
		return "", errors.New("connect grpc server fail:" + err.Error())
	}
	defer conn.Close()
	c := playerRpc.NewRpcCallClient(conn)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	res, err := c.Call(ctx, &playerRpc.RpcRequest{Body: payload})
	if err != nil {
		return "", errors.New("call remote control cmd fail:" + err.Error())
	}
	return res.Payload, nil
}

func (r RpcPlayerService) Get() (string, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", r.port), grpc.WithInsecure())
	if err != nil {
		return "", errors.New("connect grpc server fail:" + err.Error())
	}
	defer conn.Close()
	c := playerRpc.NewRpcCallClient(conn)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()
	res, err := c.Get(ctx, &playerRpc.RpcGetRequest{})
	if err != nil {
		return "", errors.New("call remote control cmd fail:" + err.Error())
	}
	//fmt.Println(res.Msg)
	return res.Payload, nil
}
