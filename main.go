package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"github.com/miguelamello/user-domain-role-service/graph"
)

// Function to convert markdown to HTML
func markdownToHTML(markdown []byte) []byte {
	html := blackfriday.Run(markdown)
	return html
}

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

func referenceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the content of the reference.md file
		content, err := os.ReadFile("./reference/reference.md")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read Reference file")
			return
		}

		// Convert Markdown to HTML
		htmlContent := markdownToHTML(content)

		// Set the appropriate headers and send the content as the response
		c.Header("Content-Type", "text/html charset=utf-8")
		c.String(http.StatusOK, string(htmlContent))
	}
}

func main() {

	// Setting up Gin in release mode
	gin.SetMode(gin.ReleaseMode)

	// Setting up Gin
	r := gin.Default()

	// API routes
	r.GET("/udrs/v1", referenceHandler())
	r.GET("/udrs/v1/playground", playgroundHandler())
	r.POST("/udrs/v1/query", graphqlHandler())
	
	r.Run("localhost:8080")

}
