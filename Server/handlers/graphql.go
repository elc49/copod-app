package handlers

import (
	"net/http"
	"time"

	"github.com/elc49/copod/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

func GraphQL() *handler.Server {
	h := handler.New(graph.NewExecutableSchema(graph.New()))
	h.AddTransport(&transport.POST{})
	// Setup websocket even though I don't need this now
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	h.Use(extension.Introspection{})
	return h
}
