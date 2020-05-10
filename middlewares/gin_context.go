package middlewares

import (
	"context"

	"golb/utils"

	"github.com/gin-gonic/gin"
)

var ginContextKey = utils.ContextKey("GinContextKey")

// WrapGinContextToContext WrapGinContextToContext
func WrapGinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GetGinContextFromContext GetGinContextFromContext
func GetGinContextFromContext(ctx context.Context) *gin.Context {
	gc, ok := ctx.Value(ginContextKey).(*gin.Context)
	if !ok {
		panic("could not retrive gin.Context")
	}
	return gc
}
