package data

import (
	"github.com/hoysics/im-kit/api/protocol"
	"log"
	"sync"
)

type Room struct {
	id   string
	mu   sync.RWMutex
	next *Channel
}

func (r *Room) Store(ch *Channel) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.next != nil {
		r.next.prev = ch
		ch.next = r.next
	}
	r.next = ch
}

func (r *Room) Delete(ch *Channel) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if ch.next != nil {
		ch.next.prev = ch.prev
	}
	if ch.prev != nil { //header point
		ch.prev.next = ch.next
	} else {
		r.next = ch.next
	}
	ch.next = nil
	ch.prev = nil
	ch.room = nil
	return
}

func (r *Room) Broadcast(major *protocol.Major) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var err error
	for n := r.next; n != nil; n = n.next {
		if !n.ValidTopic(major.GetTopic()) {
			continue
		}
		if err = n.Push(major); err != nil {
			log.Printf(`room broadcast msg to channel(%v) err:%v`, n.id, err)
		}
	}
}

func (r *Room) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for n := r.next; n != nil; n = n.next {
		n.room = nil
		n.prev = nil
		n.Close()
	}
}
