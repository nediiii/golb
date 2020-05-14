package graph

import (
	"context"
	"errors"
	"fmt"
	"golb/configs"
	"golb/graph/model"
	"golb/middlewares"
	"golb/utils"

	"github.com/99designs/gqlgen/graphql"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// [[--FUNCTION FOR DIRECTIVE]]

// HasLoginFn HasLogin impl
func HasLoginFn(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	if !configs.JwtEnable() {
		return next(ctx)
	}
	// fmt.Println("HasLoginFn被触发")

	ginContext := middlewares.GetGinContextFromContext(ctx)
	token := ginContext.Request.Header.Get("Authorization")
	if len(token) == 0 {
		return nil, errors.New("please provide a valid token in header `Authorization`")
	}
	parseToken, err := utils.JwtParse(token)
	if parseToken != nil && parseToken.Valid {
		return next(ctx)
	}
	return nil, errors.New("token invalid")
}

// HasRoleFn HasRoleFn
func HasRoleFn(ctx context.Context, obj interface{}, next graphql.Resolver, role string) (res interface{}, err error) {
	fmt.Println("HasRoleFn被触发")

	ginContext := middlewares.GetGinContextFromContext(ctx)
	token := ginContext.Request.Header.Get("Authorization")
	parseToken, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		var jwtKey = []byte("golb.sys.jwt.key")
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
		}
		return jwtKey, nil
	})
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		fmt.Println(claims["exp"], claims["iss"])
		fmt.Println("valid: ", parseToken.Valid)
		// aud exp jti iat iss nbf sub
		return next(ctx)
	}
	return nil, errors.New("token invalid")
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
