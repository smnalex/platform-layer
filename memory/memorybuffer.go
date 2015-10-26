package memory

import (
	"bytes"
	"container/list"
	log "github.com/cihub/seelog"
)

var (
	// GiveBuffer is a channel to give memory back to
	GiveBuffer chan *bytes.Buffer
	// GetBuffer is a channel to hand memory out
	GetBuffer chan *bytes.Buffer
)

func init() {
	GiveBuffer = make(chan *bytes.Buffer, channelSize)
	GetBuffer = make(chan *bytes.Buffer, channelSize)
	go bufferRecycler(GiveBuffer, GetBuffer)
}

func bufferRecycler(give, get chan *bytes.Buffer) {
	q := new(list.List)

	for {
		//This means that we are getting more than we are giving and memory is not being
		//cleaned up properly
		if q.Len() > 1000 {
			log.Warnf("Memory Recycler Overload (high memory use): %d", q.Len())
		}

		if q.Len() == 0 {
			q.PushFront(&bytes.Buffer{})
		}

		e := q.Front()

		select {
		case s := <-give:
			s.Reset()
			q.PushFront(s)

		case get <- e.Value.(*bytes.Buffer):
			q.Remove(e)
		}
	}
}
