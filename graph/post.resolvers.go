package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"strconv"
)

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *postResolver) TagConnection(ctx context.Context, obj *models.Post, first *int, last *int, after *string, before *string) (*model.PostTagsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) AuthorConnection(ctx context.Context, obj *models.Post, first *int, last *int, after *string, before *string) (*model.PostAuthorsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *postResolver) HasTags(ctx context.Context, obj *models.Post, first *int, last *int, after *string, before *string) (*model.PostTagsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *postResolver) HasAuthors(ctx context.Context, obj *models.Post, first *int, last *int, after *string, before *string) (*model.PostAuthorsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
