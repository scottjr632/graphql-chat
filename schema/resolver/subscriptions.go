package resolver

import (
	"sync"
	"time"
)

var mu sync.Mutex

var subscribers = map[string]*subscription{}

type subscriptionEvent struct {
	eventType string
	event     interface{}
}

type subscription struct {
	eventType string
	stop      <-chan struct{}
	events    chan<- interface{}
}

func atomicAddSubscriber(sub *subscription) {
	mu.Lock()
	defer mu.Unlock()
	subscribers[randomID()] = sub
}

func atomicDeleteSubscriber(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(subscribers, id)
}

func (r *Resolver) startSubscriptions() {
	unsubscribe := make(chan string)

	for {
		select {
		case id := <-unsubscribe:
			atomicDeleteSubscriber(id)
		case s := <-r.subscriptionListener:
			atomicAddSubscriber(s)
		case e := <-r.subscriptionEvent:
			// needs to be locked
			for id, s := range subscribers {
				if s.eventType == e.eventType {
					go func(id string, s *subscription) {
						select {
						case <-s.stop:
							unsubscribe <- id
							return
						default:
						}

						select {
						case <-s.stop:
							unsubscribe <- id
						case s.events <- e.event:
						case <-time.After(time.Second):
						}
					}(id, s)
				}
			}
		}
	}
}
