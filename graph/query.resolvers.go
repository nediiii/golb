package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"golb/graph/generated"
	"golb/graph/model"
	"golb/models"
	"golb/services"
	"log"
	"runtime"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SystemStatus(ctx context.Context) (*model.SysStatus, error) {
	// return nil, nil
	return &model.SysStatus{
		Arch:    runtime.GOARCH,
		Os:      runtime.GOOS,
		Version: runtime.Version(),
		NumCPU:  runtime.NumCPU(),
	}, nil
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
		return nil, errors.New("no record match")
	}
	return &v, nil
}

func (r *queryResolver) Role(ctx context.Context, id *string, name *string) (*models.Role, error) {
	tx := services.DB
	if id != nil {
		tx = tx.Where("id = ?", id)
	}
	if name != nil {
		tx = tx.Where("name = ?", name)
	}
	var v models.Role
	if tx.First(&v).RecordNotFound() {
		return nil, errors.New("no record match")
	}
	return &v, nil
}

func (r *queryResolver) User(ctx context.Context, id *string, slug *string, name *string) (*models.User, error) {
	tx := services.DB
	if id != nil {
		tx = tx.Where("id = ?", id)
	}
	if slug != nil {
		tx = tx.Where("slug = ?", slug)
	}
	if name != nil {
		tx = tx.Where("name = ?", name)
	}
	var v models.User
	if tx.First(&v).RecordNotFound() {
		return nil, errors.New("no record match")
	}
	var role models.Role
	tx.Model(&v).Related(&role, "RoleID")
	v.Role = &role
	return &v, nil
}

func (r *queryResolver) Tag(ctx context.Context, id *string, slug *string, name *string) (*models.Tag, error) {
	tx := services.DB
	if id != nil {
		tx = tx.Where("id = ?", id)
	}
	if slug != nil {
		tx = tx.Where("slug = ?", slug)
	}
	if name != nil {
		tx = tx.Where("name = ?", name)
	}
	var v models.Tag
	if tx.First(&v).RecordNotFound() {
		return nil, errors.New("no record match")
	}
	return &v, nil
}

func (r *queryResolver) Post(ctx context.Context, id *string, slug *string, name *string) (*models.Post, error) {
	tx := services.DB
	if id != nil {
		tx = tx.Where("id = ?", id)
	}
	if slug != nil {
		tx = tx.Where("slug = ?", slug)
	}
	if name != nil {
		tx = tx.Where("name = ?", name)
	}
	var v models.Post
	if tx.First(&v).RecordNotFound() {
		return nil, errors.New("no record match")
	}
	return &v, nil
}

func (r *queryResolver) AllSettings(ctx context.Context, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.SettingsConnection, error) {
	tx := services.DB
	tx = tx.Model(&models.Setting{})
	tx = tx.Order("id asc")

	// TODO use parameters to filter record

	// paging apply to tx and genertae pageInfo
	tx, pageInfo := paging(tx, *page, *perPage)

	var list []*models.Setting
	if tx.Find(&list).RecordNotFound() {
		return nil, nil
	}

	v := &model.SettingsConnection{
		Settings: list,
		PageInfo: pageInfo,
	}
	return v, nil
}

func (r *queryResolver) AllRoles(ctx context.Context, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.RolesConnection, error) {
	tx := services.DB
	tx = tx.Model(&models.Role{})
	tx = tx.Order("id asc")

	// TODO use parameters to filter record

	// paging apply to tx and genertae pageInfo
	tx, pageInfo := paging(tx, *page, *perPage)

	var list []*models.Role
	if tx.Find(&list).RecordNotFound() {
		return nil, nil
	}

	v := &model.RolesConnection{
		Roles:    list,
		PageInfo: pageInfo,
	}
	return v, nil
}

func (r *queryResolver) AllUsers(ctx context.Context, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.UsersConnection, error) {
	log.Println("AllUsers trigger")
	tx := services.DB
	tx = tx.Model(&models.User{})
	tx = tx.Order("id asc")

	// TODO use parameters to filter record

	// paging apply to tx and genertae pageInfo
	tx, pageInfo := paging(tx, *page, *perPage)

	var list []*models.User
	if tx.Find(&list).RecordNotFound() {
		return nil, nil
	}

	v := &model.UsersConnection{
		Users:    list,
		PageInfo: pageInfo,
	}
	return v, nil
}

func (r *queryResolver) AllTags(ctx context.Context, page *int, perPage *int, first *int, last *int, after *string, before *string) (*model.TagsConnection, error) {
	tx := services.DB
	tx = tx.Model(&models.Tag{})
	tx = tx.Order("id asc")

	// TODO use parameters to filter record

	// paging apply to tx and genertae pageInfo
	tx, pageInfo := paging(tx, *page, *perPage)

	var list []*models.Tag
	if tx.Find(&list).RecordNotFound() {
		return nil, nil
	}

	v := &model.TagsConnection{
		Tags:     list,
		PageInfo: pageInfo,
	}
	return v, nil
}

func (r *queryResolver) AllPosts(ctx context.Context, page *int, perPage *int, paged *bool, first *int, last *int, after *string, before *string) (*model.PostsConnection, error) {
	tx := services.DB
	tx = tx.Model(&models.Post{})
	tx = tx.Order("id asc")

	// TODO use parameters to filter record

	if paged != nil {
		tx = tx.Where("paged = ?", *paged)
	}

	// paging apply to tx and genertae pageInfo
	tx, pageInfo := paging(tx, *page, *perPage)

	var list []*models.Post
	if tx.Find(&list).RecordNotFound() {
		return nil, nil
	}

	v := &model.PostsConnection{
		Posts:    list,
		PageInfo: pageInfo,
	}
	return v, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
