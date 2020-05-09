package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/services"
)

func (r *queryResolver) Setting(ctx context.Context, id *string, key *string) (*models.Setting, error) {
	tx := services.DB
	if id != nil {
		tx = tx.Where("id = ?", id)
	}
	if key != nil {
		tx = tx.Where("key = ?", key)
	}
	var v models.Setting
	if tx.First(&v).RecordNotFound() {
		return nil, errors.New("no record")
	}
	return &v, nil
}

func (r *queryResolver) Role(ctx context.Context) (*models.Role, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tag(ctx context.Context) (*models.Tag, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllSettings(ctx context.Context, first *int, last *int, after *string, before *string) (*model.SettingsConnection, error) {
	var settings []*models.Setting
	services.DB.Model(&models.Setting{}).Find(&settings)
	v := &model.SettingsConnection{}
	v.Settings = settings
	pageInfo := &model.PageInfo{}
	pageInfo.HasNextPage = true
	v.PageInfo = pageInfo
	return v, nil
}

func (r *queryResolver) AllRoles(ctx context.Context, first *int, last *int, after *string, before *string) (*model.RolesConnection, error) {
	var list []*models.Role
	services.DB.Model(&models.Role{}).Find(&list)
	v := &model.RolesConnection{}
	v.Roles = list
	pageInfo := &model.PageInfo{}
	pageInfo.HasNextPage = true
	v.PageInfo = pageInfo
	return v, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, first *int, last *int, after *string, before *string) (*model.UsersConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllTags(ctx context.Context, first *int, last *int, after *string, before *string) (*model.TagsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPosts(ctx context.Context, first *int, last *int, after *string, before *string) (*model.PostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
