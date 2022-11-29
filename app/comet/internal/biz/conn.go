package biz

import (
	"context"
	"github.com/hoysics/im-kit/api/protocol"
	"github.com/hoysics/im-kit/app/comet/internal/data"
	"github.com/hoysics/im-kit/pkg/strings"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Conn struct {
	lr *data.LogicRepo
	cr *data.ConnRepo
	ch *data.Channel
	bk *data.Bucket

	fCtx context.Context
	cid  string
	uid  string
}

func NewConn(lr *data.LogicRepo, cr *data.ConnRepo) *Conn {
	return &Conn{lr: lr, cr: cr}
}

func (c *Conn) Connect(ctx context.Context, token string) (*data.Channel, error) {
	cid, uid, topic, msgOps, err := c.lr.Connect(ctx, token)
	if err != nil {
		return nil, err
	}
	c.cid = cid
	c.uid = uid
	c.fCtx = ctx
	//3.生成channel并存储至bucket
	mch := make(chan *protocol.Major, 5)
	c.ch = c.cr.Channel(cid, uid, mch, msgOps)
	c.bk = c.cr.Bucket(cid)
	c.bk.Store(topic, c.ch)
	return c.ch, nil
}

func (c *Conn) Operate(req *protocol.Major) error {
	var err error
	switch req.Op {
	case protocol.Operation_Heartbeat:
		//TODO count nums, if >3 or longer then 1min, return
		if err = c.lr.Heartbeat(c.fCtx, c.cid, c.uid); err != nil {
			log.Printf(`logic heartbeat err: %v`, err)
			return status.Errorf(codes.Unavailable, "ServerStreamingEcho: msg heartbeat failed")
		}
	case protocol.Operation_ChangeRoom:
		c.bk.ChangeRoom(string(req.Body), c.ch)
	case protocol.Operation_SubTopic:
		var i []int32
		if i, err = strings.SplitStrToInt32s(string(req.Body)); err != nil {
			break
		}
		c.ch.UpdateTopic(false, i)
	case protocol.Operation_UnsubTopic:
		var i []int32
		if i, err = strings.SplitStrToInt32s(string(req.Body)); err != nil {
			break
		}
		c.ch.UpdateTopic(true, i)
	case protocol.Operation_PassOnMsg:
		if err = c.lr.RecvMsg(c.fCtx, c.cid, req); err != nil {
			log.Printf(`logic recv msg err: %v`, err)
		}
	default:
		return status.Errorf(codes.NotFound, "ServerStreamingEcho: no op code matched")
	}
	return nil
}
func (c *Conn) Disconnect() {
	if c.bk != nil {
		c.bk.DeleteChannel(c.cid)
	}
	if c.lr != nil {
		c.lr.Disconnect(context.Background(), c.cid, c.uid)
	}
}
