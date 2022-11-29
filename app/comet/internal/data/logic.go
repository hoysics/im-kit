package data

import (
	"context"
	pb "github.com/hoysics/im-kit/api/logic"
	"github.com/hoysics/im-kit/api/protocol"
	"github.com/hoysics/im-kit/app/comet/internal/conf"
	"log"
)

type LogicRepo struct {
	client pb.LogicClient
}

func (r *LogicRepo) Connect(ctx context.Context, token string) (cid, uid, rid string, subMsgOps []int32, err error) {
	log.Printf(`client connect: %v`, token)
	//return "1", "1","1", []int32{1000, 1001, 1002}, nil
	connect, err := r.client.Connect(ctx, &pb.ConnectReq{
		Server: conf.ServerAddr,
		Token:  token,
	})
	if err != nil {
		return "", "", "", nil, err
	}
	return connect.GetCid(), connect.GetUserId(), connect.GetRoom(), connect.GetSubTopics(), nil
}

func (r *LogicRepo) Disconnect(ctx context.Context, cid, uid string) {
	log.Printf(`client diconnect: %v`, cid)
	if _, err := r.client.Disconnect(ctx, &pb.DisconnectReq{
		Server: conf.ServerAddr,
		Cid:    cid,
		UserId: uid,
	}); err != nil {
		return
	}
}

func (r *LogicRepo) RecvMsg(ctx context.Context, cid string, msg *protocol.Major) (err error) {
	log.Printf(`from cid(%v) recvd msg: %v`, cid, msg)
	//return
	if _, err = r.client.RecvMsg(ctx, &pb.RecvMsgReq{
		Cid: cid,
		Msg: msg,
	}); err != nil {
		return
	}
	return nil
}

func (r *LogicRepo) Heartbeat(ctx context.Context, cid, uid string) (err error) {
	log.Printf(`heartbeat: %v`, cid)
	//return
	if _, err = r.client.Heartbeat(ctx, &pb.HeartbeatReq{
		Server: conf.ServerAddr,
		Cid:    cid,
		UserId: uid,
	}); err != nil {
		return
	}
	return nil
}

func NewLogicRepo(client pb.LogicClient) *LogicRepo {
	return &LogicRepo{
		client: client,
	}
}
