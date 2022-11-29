package server

import (
	pb "github.com/hoysics/im-kit/api/comet"
	"google.golang.org/grpc"
)

func NewGRPCServer(bs pb.BrokerServer, cs pb.CometServer) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterCometServer(s, cs)
	pb.RegisterBrokerServer(s, bs)
	return s
}
