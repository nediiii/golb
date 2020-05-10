package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
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

// Resolver Resolver
type Resolver struct{}
