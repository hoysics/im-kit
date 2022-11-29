package data

import (
	"errors"
	"github.com/hoysics/im-kit/api/protocol"
	"log"
	"sync"
)

var (
	ErrMsgPipeIsFullFilled = errors.New(`channel full filled, so drop msg`)
)

type Channel struct {
	id     string
	user   string
	topics map[int32]struct{}
	pipe   chan *protocol.Major

	room *Room
	next *Channel
	prev *Channel

	mu sync.RWMutex
}

func (c *Channel) UpdateRoom(room *Room) {
	c.mu.Lock()
	c.room = room
	c.mu.Unlock()
}

func (c *Channel) Close() {
	log.Printf(`channel close, cid(%v)`, c.id)
}

func (c *Channel) ValidTopic(topic int32) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.topics[topic]; ok {
		return true
	}
	return false
}

func (c *Channel) UpdateTopic(isDel bool, topics []int32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, t := range topics {
		if isDel {
			delete(c.topics, t)
		} else {
			c.topics[t] = struct{}{}
		}
	}
	log.Printf(`cur channel(%v) tpoics: %v`, c.id, c.topics)
}

func (c *Channel) Push(m *protocol.Major) error {
	select {
	case c.pipe <- m:
	default:
		return ErrMsgPipeIsFullFilled
	}
	return nil
}

func (c *Channel) Pull() *protocol.Major {
	return <-c.pipe
}
