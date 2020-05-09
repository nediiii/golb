package dataloader

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golb/models"
	"golb/services"
)

const loadersKey = ContextKey("dataloaders")

// ContextKey ContextKey
type ContextKey string

// Loaders Loaders
type Loaders struct {
	Users RoleUsersLoader
}

// Middleware Middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			Users: RoleUsersLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(keys []uint) ([][]*models.User, []error) {
					fmt.Println("调用")
					tx := services.DB
					r := &models.Role{}
					r.ID = keys[0]
					var list []*models.User
					tx.Model(r).Related(&list, "Users").Find(&list)
					v := make([][]*models.User, 1)
					v[0] = list
					return v, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// For For
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
