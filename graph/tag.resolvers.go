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

func (r *tagResolver) ID(ctx context.Context, obj *models.Tag) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

func (r *tagResolver) PostConnection(ctx context.Context, obj *models.Tag, first *int, last *int, after *string, before *string) (*model.TagPostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *tagResolver) HasPosts(ctx context.Context, obj *models.Tag, first *int, last *int, after *string, before *string) (*model.TagPostsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
