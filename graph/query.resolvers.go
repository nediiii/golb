package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"golb/graph/generated"
	"golb/models"
	"golb/services"
)

func (r *queryResolver) Settings(ctx context.Context) ([]*models.Setting, error) {
	var settings []*models.Setting
	services.DB.Model(&models.Setting{}).Find(&settings)
	return settings, nil
}

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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
