package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"golb/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver Resolver
type Resolver struct {
	settings []*models.Setting
	setting  *models.Setting
}
