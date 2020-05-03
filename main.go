package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	"github.com/scottjr632/graphq-sub-test/graphiql"
	"github.com/scottjr632/graphq-sub-test/middleware"
	"github.com/scottjr632/graphq-sub-test/schema"
)

var httpPort = "8080"

func init() {
	port := os.Getenv("HTTP_PORT")
	if port != "" {
		httpPort = port
	}
}

func main() {
	// graphiql handler
	http.HandleFunc("/", http.HandlerFunc(graphiql.Serve(httpPort)))

	// init graphQL schema
	s, err := schema.New()
	if err != nil {
		panic(err)
	}

	// graphQL handler
	graphQLHandler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	http.HandleFunc("/graphql", middleware.LoggingHandler(graphQLHandler))

	log.Printf("Listening on port %s\n", httpPort)
	// start HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil); err != nil {
		panic(err)
	}
}
