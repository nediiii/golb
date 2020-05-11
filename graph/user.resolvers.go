package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/middlewares"
	"golb/models"
	"strconv"
)

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *userResolver) UpdateAt(ctx context.Context, obj *models.User) (string, error) {
	return strconv.FormatInt(obj.UpdatedAt.Unix(), 10), nil
}

func (r *userResolver) CreateAt(ctx context.Context, obj *models.User) (string, error) {
	return strconv.FormatInt(obj.CreatedAt.Unix(), 10), nil
}

func (r *userResolver) RoleConnection(ctx context.Context, obj *models.User, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.UserRolesConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).UserRolesLoader.Load(obj.ID)
	v := &model.UserRolesConnection{}
	v.Roles = list
	return v, nil
}

func (r *userResolver) PostConnection(ctx context.Context, obj *models.User, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.UserPostsConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).UserPostsLoader.Load(obj.ID)
	v := &model.UserPostsConnection{}
	v.Posts = list
	return v, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
