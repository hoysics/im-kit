package data

import (
	pb "github.com/hoysics/im-kit/api/logic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewLogicClient() (pb.LogicClient, func(), error) {
	conn, err := grpc.Dial(`127.0.0.1:7001`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	c := pb.NewLogicClient(conn)
	return c, func() {
		conn.Close()
	}, nil
}
