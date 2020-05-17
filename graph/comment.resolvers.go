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
	tx := services.DB
	tx = tx.Order("id desc")
	var comments []*models.Comment
	tx.Model(obj).Where("parent_id = ?", obj.ID).Find(&comments)
	v := &model.CommentRepliesConnection{}
	v.Replies = comments
	return v, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
