package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"golb/dataloader"
	"golb/graph"
	"golb/graph/generated"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.Static("/statics", "./statics")

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	log.Println("[debug] visit http://0.0.0.0:8090")
	r.Run(":8090")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	var mb int64 = 1 << 20
	h.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	return func(c *gin.Context) {
		dataloader.Middleware(h).ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {

		h.ServeHTTP(c.Writer, c.Request)
	}
}
