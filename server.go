package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"golb/graph"
	"golb/graph/generated"
	"golb/middlewares"
	"golb/services"
)

func ginSetup() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(middlewares.WrapGinContextToContext())
	r.Use(middlewares.WrapDataloaderToContext())

	r.Static("/statics", "./statics")

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	return r
}

func main() {
	services.InitDatabase()
	r := ginSetup()
	r.Run(":8090")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{
		Resolvers: &graph.Resolver{},
	}
	c.Directives.HasLogin = graph.HasLoginFn
	c.Directives.HasRole = graph.HasRoleFn

	es := generated.NewExecutableSchema(c)
	h := handler.NewDefaultServer(es)

	// add complexity limit
	complexity := &extension.ComplexityLimit{
		Func: func(ctx context.Context, rc *graphql.OperationContext) int {
			return 500
		},
	}
	h.Use(complexity)

	introspection := &extension.Introspection{}
	h.Use(introspection)

	var mb int64 = 1 << 20
	h.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {

		h.ServeHTTP(c.Writer, c.Request)
	}
}
