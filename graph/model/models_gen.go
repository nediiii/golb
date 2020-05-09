// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"golb/models"

	"github.com/99designs/gqlgen/graphql"
)

type Edges struct {
	Node   *string `json:"node"`
	Cursor *string `json:"cursor"`
}

// The `File` type, represents the response of uploading a file.
type File struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

type InputPost struct {
	ID       *string `json:"id"`
	Title    *string `json:"title"`
	HTML     *string `json:"html"`
	Markdown *string `json:"markdown"`
}

// The `InputRole` type, represents the request for set a role.
type InputRole struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// The `InputSetting` type, represents the request for set a setting.
type InputSetting struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type InputTag struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// The `InputUser` type, represents the request for set a user.
type InputUser struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Visibility *string `json:"visibility"`
	Bio        *string `json:"bio"`
}

type PageInfo struct {
	HasPreviousPage bool   `json:"hasPreviousPage"`
	HasNextPage     bool   `json:"hasNextPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type PostAuthorsConnection struct {
	PageInfo   *PageInfo          `json:"pageInfo"`
	Edges      []*PostAuthorsEdge `json:"edges"`
	TotalCount int                `json:"totalCount"`
	Authors    []*models.User     `json:"authors"`
}

type PostAuthorsEdge struct {
	Node   *models.User `json:"node"`
	Cursor string       `json:"cursor"`
}

type PostTagsConnection struct {
	PageInfo   *PageInfo       `json:"pageInfo"`
	Edges      []*PostTagsEdge `json:"edges"`
	TotalCount int             `json:"totalCount"`
	Tags       []*models.Tag   `json:"tags"`
}

type PostTagsEdge struct {
	Node   *models.Tag `json:"node"`
	Cursor string      `json:"cursor"`
}

type PostsConnection struct {
	PageInfo   *PageInfo      `json:"pageInfo"`
	Edges      []*PostsEdge   `json:"edges"`
	TotalCount int            `json:"totalCount"`
	Posts      []*models.Post `json:"posts"`
}

type PostsEdge struct {
	Node   *models.Post `json:"node"`
	Cursor string       `json:"cursor"`
}

type RoleUsersConnection struct {
	PageInfo   *PageInfo        `json:"pageInfo"`
	Edges      []*RoleUsersEdge `json:"edges"`
	TotalCount int              `json:"totalCount"`
	Users      []*models.User   `json:"users"`
}

type RoleUsersEdge struct {
	Node   *models.User `json:"node"`
	Cursor string       `json:"cursor"`
}

type RolesConnection struct {
	PageInfo   *PageInfo      `json:"pageInfo"`
	Edges      []*RolesEdge   `json:"edges"`
	TotalCount int            `json:"totalCount"`
	Roles      []*models.Role `json:"roles"`
}

type RolesEdge struct {
	Node   *models.Role `json:"node"`
	Cursor string       `json:"cursor"`
}

type SettingsConnection struct {
	PageInfo   *PageInfo         `json:"pageInfo"`
	Edges      []*SettingsEdge   `json:"edges"`
	TotalCount int               `json:"totalCount"`
	Settings   []*models.Setting `json:"settings"`
}

type SettingsEdge struct {
	Node   *models.Setting `json:"node"`
	Cursor string          `json:"cursor"`
}

type TagPostsConnection struct {
	PageInfo   *PageInfo       `json:"pageInfo"`
	Edges      []*TagPostsEdge `json:"edges"`
	TotalCount int             `json:"totalCount"`
	Posts      []*models.Post  `json:"posts"`
}

type TagPostsEdge struct {
	Node   *models.Post `json:"node"`
	Cursor string       `json:"cursor"`
}

type TagsConnection struct {
	PageInfo   *PageInfo     `json:"pageInfo"`
	Edges      []*TagsEdge   `json:"edges"`
	TotalCount int           `json:"totalCount"`
	Tags       []*models.Tag `json:"tags"`
}

type TagsEdge struct {
	Node   *models.Tag `json:"node"`
	Cursor string      `json:"cursor"`
}

// The `UploadFile` type, represents the request for uploading a file with certain payload.
type UploadFile struct {
	ID   int            `json:"id"`
	File graphql.Upload `json:"file"`
}

type UserPostsConnection struct {
	PageInfo   *PageInfo        `json:"pageInfo"`
	Edges      []*UserPostsEdge `json:"edges"`
	TotalCount int              `json:"totalCount"`
	Posts      []*models.Post   `json:"posts"`
}

type UserPostsEdge struct {
	Node   *models.Post `json:"node"`
	Cursor string       `json:"cursor"`
}

type UserRolesConnection struct {
	PageInfo   *PageInfo        `json:"pageInfo"`
	Edges      []*UserRolesEdge `json:"edges"`
	TotalCount int              `json:"totalCount"`
	Roles      []*models.Role   `json:"roles"`
}

type UserRolesEdge struct {
	Node   *models.Role `json:"node"`
	Cursor string       `json:"cursor"`
}

type UsersConnection struct {
	PageInfo   *PageInfo      `json:"pageInfo"`
	Edges      []*UsersEdge   `json:"edges"`
	TotalCount int            `json:"totalCount"`
	Users      []*models.User `json:"users"`
}

type UsersEdge struct {
	Node   *models.User `json:"node"`
	Cursor string       `json:"cursor"`
}
