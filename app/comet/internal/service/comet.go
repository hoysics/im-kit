package service

import (
	"context"
	"github.com/hoysics/im-kit/api/protocol"
	"github.com/hoysics/im-kit/app/comet/internal/biz"
	"github.com/hoysics/im-kit/app/comet/internal/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"

	pb "github.com/hoysics/im-kit/api/comet"
)

type CometService struct {
	pb.UnimplementedCometServer
	lr *data.LogicRepo
	cr *data.ConnRepo
}

func NewCometService(lr *data.LogicRepo, cr *data.ConnRepo) *CometService {
	return &CometService{
		lr: lr,
		cr: cr,
	}
}

func (s *CometService) Mail(conn pb.Comet_MailServer) (err error) {
	//1.从header取token
	md, ok := metadata.FromIncomingContext(conn.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "ServerStreamingEcho: failed to get metadata")
	}
	var token string
	if t, ok := md["token"]; ok {
		token = t[0]
	} else {
		return status.Errorf(codes.DataLoss, "ServerStreamingEcho: failed to get token in metadata")
	}
	//2.跟logic校验token 并获取用户信息
	pipe := biz.NewConn(s.lr, s.cr)
	ctx, cancel := context.WithCancel(conn.Context())
	defer func() {
		pipe.Disconnect()
		cancel()
	}()
	var ch *data.Channel
	if ch, err = pipe.Connect(ctx, token); err != nil {
		return status.Errorf(codes.Unauthenticated, "ServerStreamingEcho: invalid token")
	}
	go s.produceReply(ctx, conn, ch)
	var req *protocol.Major
	for {
		req, err = conn.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		if err := pipe.Operate(req); err != nil {
			return err
		}
		log.Printf(`recv msg: %v`, req)
	}
}

func (s *CometService) produceReply(ctx context.Context, conn pb.Comet_MailServer, ch *data.Channel) {
	for {
		select {
		case <-ctx.Done():
			log.Printf(`stream cancel`)
			return
		default:
		}
		pr := ch.Pull()
		switch pr.Op {
		case protocol.Operation_DeliverMsg:
			if err := conn.Send(pr); err != nil {
				log.Printf(`conn send err: %v`, err)
				return
			}
		default:

		}
	}
}
