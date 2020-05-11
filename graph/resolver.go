package graph

import (
	"context"
	"fmt"
	"golb/graph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jinzhu/gorm"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// [[--FUNCTION FOR DIRECTIVE]]

// HasLoginFn HasLogin impl
func HasLoginFn(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("HasLogin被触发")
	return next(ctx)
}

// HasRoleFn HasRoleFn
func HasRoleFn(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (res interface{}, err error) {
	fmt.Println("HasRoleFn被触发")
	fmt.Println("role is: ", role)
	return next(ctx)
}

// [[FUNCTION FOR DIRECTIVE--]]

// pagination
func paging(tx *gorm.DB, page, perPage int) (*gorm.DB, *model.PageInfo) {
	var count int
	tx.Count(&count)
	tx = tx.Offset((page - 1) * perPage).Limit(perPage)
	return tx, generatePageInfo(page, perPage, count)
}

// hasPreviousPage hasPreviousPage
func hasPreviousPage(page, perPage, count int) bool {
	if page == 2 && count > 0 {
		return true
	} else if page > 2 && (page-2)*perPage < count {
		return true
	}
	return false
}

// hasNextPage hasNextPage
func hasNextPage(page, perPage, count int) bool {
	if page >= 0 && page*perPage < count {
		return true
	}
	return false
}

func generatePageInfo(page, perPage, count int) *model.PageInfo {
	return &model.PageInfo{
		Total:           count,
		CurrentPage:     page,
		PerPage:         perPage,
		HasPreviousPage: hasPreviousPage(page, perPage, count),
		HasNextPage:     hasNextPage(page, perPage, count),
	}
}

// Resolver Resolver
type Resolver struct{}
