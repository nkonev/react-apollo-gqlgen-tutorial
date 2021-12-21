package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"log"
	"net/http"
	"os"
	"react-apollo-gqlgen-tutorial/backoffice/pkg/graph"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"react-apollo-gqlgen-tutorial/backoffice/pkg/middleware"
	"react-apollo-gqlgen-tutorial/backoffice/pkg/store"
	"time"
)

var (
	mb int64 = 1 << 20
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Инициализируем стор
	st := store.NewStore(store.Options{})


	srv := graph.NewServer(graph.Options{
		Store: st,
	})
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: transport.WebsocketInitFunc(func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return ctx, nil
		}),
	})
	srv.Use(extension.Introspection{})

	// Создадим роутер
	router := mux.NewRouter()

	// Инициализируем middleware
	// Передадим Store в качестве параметра
	router.Use(middleware.AuthMiddleware(st))
	router.Use(middleware.CorsMiddleware(st))

	router.Handle("/", playground.Handler("GraphQL playground", "/graph"))
	router.Handle("/graph", srv)


	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
