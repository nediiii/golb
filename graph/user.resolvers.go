package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/middleware"
	"golb/models"
	"strconv"
)

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *userResolver) RoleConnection(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserRolesConnection, error) {
	list, _ := middleware.GetDataloaderFromContext(ctx).UserRolesLoader.Load(obj.ID)
	v := &model.UserRolesConnection{}
	v.Roles = list
	return v, nil
}

func (r *userResolver) PostConnection(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserPostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
