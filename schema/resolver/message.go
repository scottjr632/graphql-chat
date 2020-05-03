package resolver

import (
	"context"
	"sync"
)

var mutex sync.Mutex

var messages = []*Message{
	&Message{id: randomID(), content: "test", channel: &Channel{
		id:   randomID(),
		name: "test",
	}},
}

// Message is the GraphQL message object type
type Message struct {
	id      string
	content string
	channel *Channel
}

func (r *Resolver) MessageCreated(ctx context.Context, arg struct{ ChannelName string }) <-chan *Message {
	c := make(chan interface{})
	mc := make(chan *Message)

	r.subscriptionListener <- &subscription{
		eventType: "message",
		events:    c,
		stop:      ctx.Done(),
	}

	// casts the interface type to Message pointer type
	go func() {
		for {
			select {
			case d := <-c:
				m := d.(*Message)
				if m.Channel().Name() == arg.ChannelName {
					mc <- d.(*Message)
				}
			}
		}
	}()

	return mc
}

// CreateMessage is the message mutation
func (r *Resolver) CreateMessage(args struct {
	Content     string
	ChannelName string
}) *Message {
	newMessage := &Message{
		id:      randomID(),
		content: args.Content,
		channel: &Channel{
			id:   randomID(),
			name: args.ChannelName,
		},
	}

	mutex.Lock()
	defer mutex.Unlock()
	messages = append(messages, newMessage)
	r.subscriptionEvent <- &subscriptionEvent{
		eventType: "message",
		event:     newMessage,
	}
	return newMessage
}

// Messages resolves for the messages query
func (r *Resolver) Messages() *[]*Message {
	return &messages
}

// ID is the id for the messages
func (m *Message) ID() string {
	return m.id
}

// Content ...
func (m *Message) Content() string {
	return m.content
}

// Channel ...
func (m *Message) Channel() *Channel {
	return m.channel
}
