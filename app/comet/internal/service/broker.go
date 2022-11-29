package service

import (
	"context"
	"github.com/hoysics/im-kit/app/comet/internal/data"
	"log"

	pb "github.com/hoysics/im-kit/api/comet"
)

type BrokerService struct {
	pb.UnimplementedBrokerServer

	cr *data.ConnRepo
}

func NewBrokerService(cr *data.ConnRepo) *BrokerService {
	return &BrokerService{
		cr: cr,
	}
}

func (s *BrokerService) PushMsg(ctx context.Context, req *pb.PushMsgRequest) (*pb.PushMsgReply, error) {
	var bk *data.Bucket
	for _, cid := range req.GetCids() {
		if bk = s.cr.Bucket(cid); bk == nil {
			continue
		}
		bk.PushMsg(cid, req.GetMsg())
	}
	log.Printf(`finish push msg: %v`, req)
	return &pb.PushMsgReply{}, nil
}
func (s *BrokerService) Broadcast(ctx context.Context, req *pb.BroadcastRequest) (*pb.BroadcastReply, error) {
	s.cr.TraversalBuckets(func(bucket *data.Bucket) {
		//TODO use speed to control
		bucket.Broadcast(req.GetSpeed(), req.GetMsg())
	})
	log.Printf(`finish broadcast: %v`, req)
	return &pb.BroadcastReply{}, nil
}
func (s *BrokerService) BroadcastInRoom(ctx context.Context, req *pb.BroadcastInRoomRequest) (*pb.BroadcastInRoomReply, error) {
	s.cr.TraversalBuckets(func(bucket *data.Bucket) {
		bucket.BroadcastInRoom(req.GetRoom(), req.GetMsg())
	})
	log.Printf(`finish broadcast in room: %v`, req)
	return &pb.BroadcastInRoomReply{}, nil
}
