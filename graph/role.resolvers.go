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

func (r *roleResolver) ID(ctx context.Context, obj *models.Role) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *roleResolver) UpdateAt(ctx context.Context, obj *models.Role) (string, error) {
	return strconv.FormatInt(obj.UpdatedAt.Unix(), 10), nil
}

func (r *roleResolver) CreateAt(ctx context.Context, obj *models.Role) (string, error) {
	return strconv.FormatInt(obj.CreatedAt.Unix(), 10), nil
}

func (r *roleResolver) UserConnection(ctx context.Context, obj *models.Role, first *int, last *int, after *string, before *string) (*model.RoleUsersConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).RoleUsersLoader.Load(obj.ID)
	v := &model.RoleUsersConnection{}
	v.Users = list
	return v, nil
}

// Role returns generated.RoleResolver implementation.
func (r *Resolver) Role() generated.RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }
