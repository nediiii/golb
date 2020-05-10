package dataloader

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golb/models"
	"golb/services"

	"github.com/jinzhu/gorm"
)

const loadersKey = ContextKey("dataloaders")

// ContextKey ContextKey
type ContextKey string

// DataLoader DataLoader
type DataLoader interface {
}

// Loaders Loaders
type Loaders struct {
	RoleUsersLoader
	UserRolesLoader
}

// loaders loaders
var (
	loaders = &Loaders{
		RoleUsersLoader: RoleUsersLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.User, []error) {
				fmt.Println("调用")
				tx := services.DB
				v := make([][]*models.User, len(keys))
				for i, k := range keys {
					tx.Model(&models.Role{Model: gorm.Model{ID: k}}).Related(&v[i], "Users")
				}
				return v, nil
			},
		},
		UserRolesLoader: UserRolesLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.Role, []error) {
				fmt.Println("调用")
				tx := services.DB
				v := make([][]*models.Role, len(keys))
				for i, k := range keys {
					tx.Model(&models.User{Model: gorm.Model{ID: k}}).Related(&v[i], "Roles")
				}
				return v, nil
			},
		},
	}
)

// Middleware Middleware
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

// For For
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
