package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/miguelamello/user-domain-role-service/graph"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL", "/udrs/v1/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func main() {

	// Setting up Gin
	r := gin.Default()

	// API routes
	r.POST("/udrs/v1/query", graphqlHandler())
	r.GET("/udrs/v1", playgroundHandler())
	r.Run("localhost:8080")

}
