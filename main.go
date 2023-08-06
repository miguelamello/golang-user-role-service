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

	h := playground.Handler("GraphQL", "/udrs/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func main() {

	// Setting up Gin
	r := gin.Default()
	r.POST("/udrs/query", graphqlHandler())
	r.GET("/udrs", playgroundHandler())
	r.Run()

}
