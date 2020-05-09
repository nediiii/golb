package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"golb/graph/generated"
	"golb/models"
	"golb/services"

	"github.com/jinzhu/gorm"
)

func (r *mutationResolver) CreateSetting(ctx context.Context, key string, value string) (*models.Setting, error) {
	s := &models.Setting{}
	s.Key = key
	s.Value = value
	var err gorm.Errors
	err = services.DB.Create(s).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return s, nil
}

func (r *mutationResolver) DeleteSetting(ctx context.Context, id string) (bool, error) {
	var err gorm.Errors
	err = services.DB.Where("id = ?", id).Delete(models.Setting{}).GetErrors()
	if len(err) > 0 {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateSetting(ctx context.Context, id string, key string, value string) (*models.Setting, error) {
	v := &models.Setting{}
	tx := services.DB
	if tx.First(v, id).RecordNotFound() {
		return nil, errors.New("要更新的记录不存在")
	}
	v.Key = key
	v.Value = value
	var err gorm.Errors
	tx.Save(v).GetErrors()
	if len(err) > 0 {
		return nil, err
	}
	return v, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
