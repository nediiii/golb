package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golb/dataloader"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"strconv"
)

func (r *roleResolver) ID(ctx context.Context, obj *models.Role) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *roleResolver) UserConnection(ctx context.Context, obj *models.Role, first *int, last *int, after *string, before *string) (*model.RoleUsersConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Role returns generated.RoleResolver implementation.
func (r *Resolver) Role() generated.RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *roleResolver) HasUsers(ctx context.Context, obj *models.Role, first *int, last *int, after *string, before *string) (*model.RoleUsersConnection, error) {
	users, _ := dataloader.For(ctx).Users.Load(obj.ID)
	v := &model.RoleUsersConnection{}
	v.Users = users
	return v, nil
}
