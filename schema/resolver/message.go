package resolver

import (
	"context"
	"sync"

	"github.com/scottjr632/graphql-chat/routes"
)

var mutex sync.Mutex

var messages = []*Message{
	&Message{id: randomID(), content: "Hello, world!", channel: &Channel{
		id:   randomID(),
		name: "Memes",
	}},
}

// Message is the GraphQL message object type
type Message struct {
	id      string
	content string
	channel *Channel
}

func (m *Message) toJSON() *messageJSON {
	return &messageJSON{
		ID:      m.ID(),
		Content: m.Content(),
		Channel: m.Channel().toJSON(),
	}
}

type messageJSON struct {
	ID      string       `json:"id"`
	Content string       `json:"content"`
	Channel *ChannelJSON `json:"channel"`
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

	routes.GetSubscriptionService().Send(newMessage.toJSON())
	return newMessage
}

// Messages resolves for the messages query
func (r *Resolver) Messages(args struct{ ChannelName string }) *[]*Message {
	var channelMessages []*Message
	for _, message := range messages {
		if message.channel.name == args.ChannelName {
			channelMessages = append(channelMessages, message)
		}
	}
	return &channelMessages
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
