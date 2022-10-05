package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gregkolb/example/graph"
	"github.com/gregkolb/example/graph/generated"
	"github.com/rs/cors"
	// "github.com/gorilla/websocket"
)

const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

// 	srv.AddTransport(&transport.Websocket{}) // <---- This is the important part!

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000", "*"},
	// 	AllowCredentials: true,
	// 	Debug:            true,
	// }).Handler)

	router.Use(cors.Default().Handler)
	
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{})
    // srv.AddTransport(&transport.Websocket{
    //     Upgrader: websocket.Upgrader{
    //         CheckOrigin: func(r *http.Request) bool {
    //             // Check against your desired domains here
    //              return true
    //         },
    //         ReadBufferSize:  1024,
    //         WriteBufferSize: 1024,
    //     },
    // })

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}