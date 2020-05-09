package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golb/graph/generated"
	"golb/models"
	"strconv"
)

func (r *settingResolver) ID(ctx context.Context, obj *models.Setting) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
}

// Setting returns generated.SettingResolver implementation.
func (r *Resolver) Setting() generated.SettingResolver { return &settingResolver{r} }

type settingResolver struct{ *Resolver }
