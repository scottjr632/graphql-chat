package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"github.com/scottjr632/graphql-chat/graphiql"
	"github.com/scottjr632/graphql-chat/middleware"
	"github.com/scottjr632/graphql-chat/schema"
)

var httpPort = "8080"

func init() {
	port := os.Getenv("HTTP_PORT")
	if port != "" {
		httpPort = port
	}
}

func ginHTTPAdapter(f func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	r.GET("/", ginHTTPAdapter(http.HandlerFunc(graphiql.Serve(httpPort))))

	s, err := schema.New()
	if err != nil {
		panic(err)
	}

	graphQLHandler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	r.Any("/graphql", ginHTTPAdapter(middleware.CORS(middleware.LoggingHandler(graphQLHandler))))

	r.Run()
}
