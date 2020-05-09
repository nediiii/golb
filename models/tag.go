package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Tag tag model
type Tag struct {
	gorm.Model
	Posts           []*Post   `gorm:"many2many:post_tags"`
	UUID            uuid.UUID `gorm:"type:uuid;unique_index"`
	Name            string    `gorm:"unique_index;not null"`
	Slug            string
	Description     string
	Image           string
	Visibility      string
	MetaTitle       string
	MetaDescription string
	CreateBy        uint
	UpdateBy        uint
}


// GetID GetID
func (v Tag) GetID() interface{} {
	return v.ID
}

// GetCreateAt GetCreateAt
func (v Tag) GetCreateAt() interface{} {
	return v.CreatedAt
}

// GetUpdateAt GetUpdateAt
func (v Tag) GetUpdateAt() interface{} {
	return v.UpdatedAt
}


// BeforeCreate 初始化uuid
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedTags PreDefinedTags
var PreDefinedTags = []*Tag{
	{Name: "article", Description: "博客文章"},
	{Name: "learn", Description: "学习笔记"},
	{Name: "goweb", Description: "动态网站"},
	{Name: "python", Description: "python"},
	{Name: "javascript", Description: "javascript"},
}
