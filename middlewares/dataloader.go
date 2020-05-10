package middlewares

import (
	"context"
	"golb/dataloader"
	"golb/utils"

	"github.com/gin-gonic/gin"
)

const loadersKey = utils.ContextKey("dataloaders")

// WrapDataloaderToContext WrapDataloaderToContext
func WrapDataloaderToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, dataloader.Loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GetDataloaderFromContext GetDataloaderFromContext
func GetDataloaderFromContext(ctx context.Context) *dataloader.DataLoaders {
	loaders, ok := ctx.Value(loadersKey).(*dataloader.DataLoaders)
	if !ok {
		panic("could not retrive dataloader")
	}
	return loaders
}
