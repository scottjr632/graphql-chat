package resolver

import "time"

type HelloSaidEvent struct {
	id  string
	msg string
}

type helloSaidSubscriber struct {
	stop   <-chan struct{}
	events chan<- *HelloSaidEvent
}

func (r *HelloSaidEvent) Msg() string {
	return r.msg
}

func (r *HelloSaidEvent) ID() string {
	return r.id
}

func (r *Resolver) broadcastHelloSaid() {
	subscribers := map[string]*helloSaidSubscriber{}
	unsubscribe := make(chan string)

	// NOTE: subscribing and sending events are at odds.
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.helloSaidSubscriber:
			subscribers[randomID()] = s
		case e := <-r.helloSaidEvents:
			for id, s := range subscribers {
				go func(id string, s *helloSaidSubscriber) {
					select {
					case <-s.stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.stop:
						unsubscribe <- id
					case s.events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}
		}
	}
}
