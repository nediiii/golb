package dataloader

import (
	"fmt"
	"time"

	"golb/models"
	"golb/services"

	"github.com/jinzhu/gorm"
)

// DataLoaders DataLoaders
type DataLoaders struct {
	RoleUsersLoader
	UserRolesLoader
}

// Loaders Loaders
var (
	Loaders = &DataLoaders{
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
