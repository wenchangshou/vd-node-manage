package rpcServer

import (
	"context"
	"github.com/wenchangshou2/vd-node-manage/common/publisher"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/rpc/server/pb"
	"strings"
	"time"
)

var (
	G_pubsubSerice *PubSubServer
)

type PubSubServer struct {
	pub *publisher.Publisher
	pb.UnimplementedPubsubServiceServer
}

func NewPubsService() *PubSubServer {
	G_pubsubSerice = &PubSubServer{
		pub: publisher.NewPublisher(100*time.Millisecond, 10),
	}
	return G_pubsubSerice
}
func (p *PubSubServer) Publish(ctx context.Context, req *pb.PublishChannel) (*pb.Empty, error) {
	p.pub.Publish(req)
	return &pb.Empty{}, nil
}
func (p *PubSubServer) Subscribe(channel *pb.SubscribeChannel, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if rev, ok := v.(*pb.PublishChannel); ok {
			return strings.HasPrefix(rev.Topic,channel.Topic)
		}
		return false
	})
	for v := range ch {
		tmp := v.(*pb.PublishChannel)
		if err := stream.Send(&pb.SubscribeResult{
			Id:     tmp.Id,
			Topic:  tmp.Topic,
			Action: tmp.Action,
			Body:   tmp.Body,
		}); err != nil {
			return err
		}
	}
	return nil
}
