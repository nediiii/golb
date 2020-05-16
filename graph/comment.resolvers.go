package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/utils"
)

func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (string, error) {
	return utils.Uint2String(obj.ID), nil
}

func (r *commentResolver) UpdateAt(ctx context.Context, obj *models.Comment) (string, error) {
	return utils.Time2UnixString(&obj.UpdatedAt), nil
}

func (r *commentResolver) CreateAt(ctx context.Context, obj *models.Comment) (string, error) {
	return utils.Time2UnixString(&obj.CreatedAt), nil
}

func (r *commentResolver) Post(ctx context.Context, obj *models.Comment) (*models.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Parent(ctx context.Context, obj *models.Comment) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) ReplyConnection(ctx context.Context, obj *models.Comment, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.CommentRepliesConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *commentResolver) PostID(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *commentResolver) ParentID(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
