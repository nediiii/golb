package main

import (
	"context"
	"os"
	"strings"

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
	var r *gin.Engine

	if strings.HasSuffix(os.Args[0], ".test") {
		// turn off log
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	} else {
		r = gin.Default()
	}

	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.Use(middlewares.WrapGinContextToContext())
	r.Use(middlewares.WrapDataloaderToContext())

	if !strings.HasSuffix(os.Args[0], ".test") {
		r.Static("/statics", "./statics")
		r.GET("/", playgroundHandler())
	}

	r.POST("/query", graphqlHandler())
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

	// 1 2 4 8
	// scalar type connection

	connectionComplexityFn := func(childComplexity int, page, perPage, first, last *int, after, before *string) int {
		return childComplexity * 4
	}
	c.Complexity.Role.UserConnection = connectionComplexityFn
	c.Complexity.User.PostConnection = connectionComplexityFn
	c.Complexity.Post.AuthorConnection = connectionComplexityFn
	c.Complexity.Post.CommentConnection = connectionComplexityFn
	c.Complexity.Post.TagConnection = connectionComplexityFn
	c.Complexity.Tag.PostConnection = connectionComplexityFn
	c.Complexity.Comment.ReplyConnection = connectionComplexityFn

	es := generated.NewExecutableSchema(c)
	h := handler.NewDefaultServer(es)

	// add complexity limit
	complexity := &extension.ComplexityLimit{
		Func: func(ctx context.Context, rc *graphql.OperationContext) int {
			return 1 << 9 // 512
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
