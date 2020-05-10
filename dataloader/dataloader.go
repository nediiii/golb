package dataloader

import (
	"time"

	"golb/models"
	"golb/services"

	"github.com/jinzhu/gorm"
)

// cd dataloader (!! need cd to current directory )
// go run github.com/vektah/dataloaden RoleUsersLoader uint []*golb/models.User
// go run github.com/vektah/dataloaden UserRolesLoader uint []*golb/models.Role
// go run github.com/vektah/dataloaden UserPostsLoader uint []*golb/models.Post
// go run github.com/vektah/dataloaden PostTagsLoader uint []*golb/models.Tag
// go run github.com/vektah/dataloaden PostAuthorsLoader uint []*golb/models.User
// go run github.com/vektah/dataloaden TagPostLoader uint []*golb/models.Post

// DataLoaders DataLoaders
type DataLoaders struct {
	RoleUsersLoader
	UserRolesLoader
	UserPostsLoader
	PostTagsLoader
	PostAuthorsLoader
	TagPostLoader
}

// Loaders Loaders
var (
	Loaders = &DataLoaders{
		RoleUsersLoader: RoleUsersLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.User, []error) {
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
				tx := services.DB
				v := make([][]*models.Role, len(keys))
				for i, k := range keys {
					tx.Model(&models.User{Model: gorm.Model{ID: k}}).Related(&v[i], "Roles")
				}
				return v, nil
			},
		},
		UserPostsLoader: UserPostsLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.Post, []error) {
				tx := services.DB
				v := make([][]*models.Post, len(keys))
				for i, k := range keys {
					tx.Model(&models.User{Model: gorm.Model{ID: k}}).Related(&v[i], "Posts")
				}
				return v, nil
			},
		},
		PostTagsLoader: PostTagsLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.Tag, []error) {
				tx := services.DB
				v := make([][]*models.Tag, len(keys))
				for i, k := range keys {
					tx.Model(&models.Post{Model: gorm.Model{ID: k}}).Related(&v[i], "Tags")
				}
				return v, nil
			},
		},
		PostAuthorsLoader: PostAuthorsLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.User, []error) {
				tx := services.DB
				v := make([][]*models.User, len(keys))
				for i, k := range keys {
					tx.Model(&models.Post{Model: gorm.Model{ID: k}}).Related(&v[i], "Users")
				}
				return v, nil
			},
		},
		TagPostLoader: TagPostLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []uint) ([][]*models.Post, []error) {
				tx := services.DB
				v := make([][]*models.Post, len(keys))
				for i, k := range keys {
					tx.Model(&models.Tag{Model: gorm.Model{ID: k}}).Related(&v[i], "Posts")
				}
				return v, nil
			},
		},
	}
)
