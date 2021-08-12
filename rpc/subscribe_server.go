package rpc

import (
	"context"
	"strings"
	"time"

	"github.com/wenchangshou2/vd-node-manage/pkg/pubsub"
	"github.com/wenchangshou2/vd-node-manage/rpc/pb"
)

var (
	G_pubsubSerice *PubSubServer
)

type PubSubServer struct {
	pub *pubsub.Publisher
	pb.UnimplementedPubsubServiceServer
}

func NewPubsubService() *PubSubServer {
	G_pubsubSerice = &PubSubServer{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
	return G_pubsubSerice
}
func (p *PubSubServer) Publish(ctx context.Context, req *pb.PublishChannel) (*pb.Empty, error) {
	p.pub.Publish(req)
	return &pb.Empty{}, nil
}
func (p *PubSubServer) Subscribe(channel *pb.SubscribeChannel, stream pb.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if recv, ok := v.(*pb.PublishChannel); ok {
			if strings.HasPrefix(recv.Topic, channel.Topic) {
				return true
			}
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
