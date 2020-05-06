package resolver

import (
	"math/rand"
)

// Resolver is the base query resolver for GraphQL
type Resolver struct {
	subscriptionListener chan *subscription
	subscriptionEvent    chan *subscriptionEvent
}

// New returns a new resolver
func New() *Resolver {
	r := &Resolver{
		subscriptionListener: make(chan *subscription),
		subscriptionEvent:    make(chan *subscriptionEvent),
	}

	go r.startSubscriptions()

	return r
}

// Hello is resolves for the hello query
func (r *Resolver) Hello() string {
	return "Hello world!"
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
