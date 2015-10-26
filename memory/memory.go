package memory

import (
	"container/list"
	log "github.com/cihub/seelog"
)

const (
	defaultMemorySize = 100
	channelSize       = 10
)

var (
	// Give is a channel to give memory back to
	Give chan []byte
	// Get is a channel to hand memory out
	Get chan []byte
)

func init() {
	Give = make(chan []byte, channelSize)
	Get = make(chan []byte, channelSize)
	go recycler(Give, Get)
}

func recycler(give, get chan []byte) {
	q := new(list.List)

	for {
		//This means that we are getting more than we are giving and memory is not being
		//cleaned up properly
		if q.Len() > 1000 {
			log.Warnf("Memory Recycler Overload (high memory use): %d", q.Len())
		}

		if q.Len() == 0 {
			q.PushFront(make([]byte, defaultMemorySize))
		}

		e := q.Front()

		select {
		case s := <-give:
			q.PushFront(s[:0])

		case get <- e.Value.([]byte):
			q.Remove(e)
		}
	}
}
