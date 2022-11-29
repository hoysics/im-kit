package data

import (
	"github.com/hoysics/im-kit/api/protocol"
	"github.com/hoysics/im-kit/app/comet/internal/conf"
	"hash/adler32"
	"sync"
)

type ConnRepo struct {
	chs map[string]*Channel
	bks []*Bucket
	mu  sync.RWMutex
}

func (r *ConnRepo) Channel(cid, uid string, ch chan *protocol.Major, topics []int32) *Channel {
	c := &Channel{
		id:     cid,
		user:   uid,
		pipe:   ch,
		topics: make(map[int32]struct{}),
	}
	for _, t := range topics {
		c.topics[t] = struct{}{}
	}
	return c
}

func (r *ConnRepo) Bucket(cid string) *Bucket {
	h := adler32.New()
	if _, err := h.Write([]byte(cid)); err != nil {
		return nil
	}
	c := h.Sum32() % conf.BucketSize
	return r.bks[c]
}

func (r *ConnRepo) TraversalBuckets(fn func(bucket *Bucket)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, bk := range r.bks {
		fn(bk)
	}
}

func NewConnRepo() *ConnRepo {
	r := &ConnRepo{
		chs: make(map[string]*Channel),
		bks: make([]*Bucket, conf.BucketSize),
	}
	for i := 0; i < conf.BucketSize; i++ {
		r.bks[i] = &Bucket{
			mu:    sync.RWMutex{},
			chs:   make(map[string]*Channel),
			rooms: make(map[string]*Room),
		}
	}
	return r
}
