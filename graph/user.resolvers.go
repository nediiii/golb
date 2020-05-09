package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/services"
	"strconv"
)

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *userResolver) RoleConnection(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserRolesConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) PostConnection(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserPostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) HasRoles(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserRolesConnection, error) {
	v := &model.UserRolesConnection{}
	var list []*models.Role
	services.DB.Model(obj).Related(&list, "Roles").Find(&list)
	v.Roles = list
	return v, nil
}
func (r *userResolver) HasPosts(ctx context.Context, obj *models.User, first *int, last *int, after *string, before *string) (*model.UserPostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
