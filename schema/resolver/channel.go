package resolver

import "sync"

var channelMu sync.Mutex

var channels = []*Channel{
	&Channel{
		id:   randomID(),
		name: "Memes",
	},
	&Channel{
		id:   randomID(),
		name: "Random",
	},
}

// Channel represents the Channel GraphQL object type
type Channel struct {
	id   string
	name string
}

func (r *Resolver) CreateChannel(arg struct{ Name string }) *Channel {
	newChannel := &Channel{id: randomID(), name: arg.Name}
	mu.Lock()
	defer mu.Unlock()
	channels = append(channels, newChannel)
	return newChannel
}

func (r *Resolver) Channels() *[]*Channel {
	return &channels
}

func (r *Resolver) Channel(arg struct{ Name string }) *Channel {
	for _, channel := range channels {
		if channel.name == arg.Name {
			return channel
		}
	}
	return nil
}

// ID ...
func (c *Channel) ID() string {
	return c.id
}

// Name ...
func (c *Channel) Name() string {
	return c.name
}
