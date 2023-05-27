package main

import (
	"log"
	"net/http"
	"os"
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/miguelamello/graph"
)

const defaultPort = "8080"

type Resolver struct{}

type queryResolver struct{}

func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{}
}



func (q *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
