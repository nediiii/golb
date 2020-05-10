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

func (r *tagResolver) ID(ctx context.Context, obj *models.Tag) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *tagResolver) UpdateAt(ctx context.Context, obj *models.Tag) (string, error) {
	return strconv.FormatInt(obj.UpdatedAt.Unix(), 10), nil
}

func (r *tagResolver) CreateAt(ctx context.Context, obj *models.Tag) (string, error) {
	return strconv.FormatInt(obj.CreatedAt.Unix(), 10), nil
}

func (r *tagResolver) PostConnection(ctx context.Context, obj *models.Tag, first *int, last *int, after *string, before *string) (*model.TagPostsConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).TagPostLoader.Load(obj.ID)
	v := &model.TagPostsConnection{}
	v.Posts = list
	return v, nil
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
