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

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *postResolver) UpdateAt(ctx context.Context, obj *models.Post) (string, error) {
	return strconv.FormatInt(obj.UpdatedAt.Unix(), 10), nil
}

func (r *postResolver) CreateAt(ctx context.Context, obj *models.Post) (string, error) {
	return strconv.FormatInt(obj.CreatedAt.Unix(), 10), nil
}

func (r *postResolver) TagConnection(ctx context.Context, obj *models.Post, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.PostTagsConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).PostTagsLoader.Load(obj.ID)
	v := &model.PostTagsConnection{}
	v.Tags = list
	return v, nil
}

func (r *postResolver) AuthorConnection(ctx context.Context, obj *models.Post, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.PostAuthorsConnection, error) {
	list, _ := middlewares.GetDataloaderFromContext(ctx).PostAuthorsLoader.Load(obj.ID)
	v := &model.PostAuthorsConnection{}
	v.Authors = list
	return v, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
