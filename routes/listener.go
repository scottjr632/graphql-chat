package routes

import (
	"sync"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

var listenerOnce sync.Once
var subscriptionServiceInstance *subscriptionService

type subscriptionService struct {
	channel chan interface{}
}

// SubscriptionService holds the Send and Receive for server events
type SubscriptionService interface {
	Send(interface{})
	Receive() chan interface{}
}

func (s *subscriptionService) Send(data interface{}) {
	s.channel <- data
}

func (s *subscriptionService) Receive() chan interface{} {
	return s.channel
}

// GetSubscriptionService returns a single service
func GetSubscriptionService() SubscriptionService {
	listenerOnce.Do(func() {
		subscriptionServiceInstance = &subscriptionService{
			channel: make(chan interface{}),
		}
	})
	return subscriptionServiceInstance
}

// Subscribe allows a user to subscribe to a server events
func (r *Routes) Subscribe(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")

	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-c.Request.Context().Done():
			// release resources
			break
		case data := <-GetSubscriptionService().Receive():
			sse.Encode(c.Writer, sse.Event{
				Event: "message",
				Data:  data,
			})

			c.Writer.Flush()
		case <-ticker.C:
			sse.Encode(c.Writer, sse.Event{
				Event: "heartbeat",
				Data:  "Alive",
			})

			c.Writer.Flush()
		}
	}
}
