package resolver

import (
	"context"
	"math/rand"
	"time"
)

// Resolver is the base query resolver for GraphQL
type Resolver struct {
	helloSaidEvents      chan *HelloSaidEvent
	helloSaidSubscriber  chan *helloSaidSubscriber
	subscriptionListener chan *subscription
	subscriptionEvent    chan *subscriptionEvent
}

// New returns a new resolver
func New() *Resolver {
	r := &Resolver{
		helloSaidEvents:     make(chan *HelloSaidEvent),
		helloSaidSubscriber: make(chan *helloSaidSubscriber),

		subscriptionListener: make(chan *subscription),
		subscriptionEvent:    make(chan *subscriptionEvent),
	}

	go r.broadcastHelloSaid()
	go r.startSubscriptions()

	return r
}

// Hello is resolves for the hello query
func (r *Resolver) Hello() string {
	return "Hello world!"
}

// SayHello is a mutation
func (r *Resolver) SayHello(args struct{ Msg string }) *HelloSaidEvent {
	e := &HelloSaidEvent{msg: args.Msg, id: randomID()}
	go func() {
		select {
		case r.helloSaidEvents <- e:
		case <-time.After(1 * time.Second):
		}
	}()
	return e
}

// HelloSaid is the subscriptoip
func (r *Resolver) HelloSaid(ctx context.Context) <-chan *HelloSaidEvent {
	c := make(chan *HelloSaidEvent)
	// NOTE: this could take a while
	r.helloSaidSubscriber <- &helloSaidSubscriber{events: c, stop: ctx.Done()}

	return c
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
