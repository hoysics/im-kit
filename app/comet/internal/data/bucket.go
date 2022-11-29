package data

import (
	"github.com/hoysics/im-kit/api/protocol"
	"log"
	"sync"
)

type Bucket struct {
	mu    sync.RWMutex
	chs   map[string]*Channel
	rooms map[string]*Room
}

func (b *Bucket) Store(rid string, ch *Channel) {
	b.mu.Lock()
	if ch, ok := b.chs[ch.id]; ok {
		ch.Close()
	}
	b.chs[ch.id] = ch
	b.mu.Unlock()
	var (
		r  *Room
		ok bool
	)
	if rid != "" {
		b.mu.Lock()
		if r, ok = b.rooms[rid]; !ok {
			r = &Room{}
			b.rooms[rid] = r
		}
		b.mu.Unlock()
	}
	if r != nil {
		r.Store(ch)
		ch.UpdateRoom(r)
	}
}

func (b *Bucket) Load(cid string) *Channel {
	b.mu.RLock()
	ch, ok := b.chs[cid]
	b.mu.RUnlock()
	if ok {
		return ch
	}
	return nil
}

func (b *Bucket) DeleteChannel(cid string) {
	var ch *Channel
	var ok bool
	//1. del in chs
	b.mu.Lock()
	if ch, ok = b.chs[cid]; ok {
		delete(b.chs, cid)
	}
	b.mu.Unlock()
	//2. del in room
	if ch == nil {
		return
	}
	if r := ch.room; r != nil {
		r.Delete(ch)
	}
}

func (b *Bucket) DeleteRoom(rid string) {
	var r *Room
	var ok bool
	b.mu.Lock()
	if r, ok = b.rooms[rid]; ok {
		delete(b.rooms, rid)
	}
	b.mu.Unlock()
	r.Close()
}

func (b *Bucket) ChangeRoom(nrid string, ch *Channel) {
	var (
		r  *Room
		ok bool
	)
	if or := ch.room; or != nil {
		or.Delete(ch)
	}
	if nrid != "" {
		b.mu.Lock()
		if r, ok = b.rooms[nrid]; !ok {
			r = &Room{}
			b.rooms[nrid] = r
		}
		b.mu.Unlock()
	}
	if r != nil {
		r.Store(ch)
		ch.UpdateRoom(r)
	}
}

func (b *Bucket) PushMsg(cid string, major *protocol.Major) {
	var ch *Channel
	var ok bool
	b.mu.RLock()
	ch, ok = b.chs[cid]
	b.mu.RUnlock()
	if !ok {
		return
	}
	if !ch.ValidTopic(major.GetTopic()) {
		return
	}
	if err := ch.Push(major); err != nil {
		log.Printf(`bucket push msg to channel(%v) err:%v`, ch.id, err)
	}
}

func (b *Bucket) BroadcastInRoom(rid string, major *protocol.Major) {
	var r *Room
	var ok bool
	b.mu.RLock()
	r, ok = b.rooms[rid]
	b.mu.RUnlock()
	if !ok {
		return
	}
	r.Broadcast(major)

}

func (b *Bucket) Broadcast(speed int32, major *protocol.Major) {
	var ch *Channel
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, ch = range b.chs {
		if !ch.ValidTopic(major.GetTopic()) {
			continue
		}
		if err := ch.Push(major); err != nil {
			log.Printf(`bucket push msg to channel(%v) err:%v`, ch.id, err)
		}
	}
}
