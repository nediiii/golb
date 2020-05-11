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
	Slug            string    `gorm:"unique_index;not null"`
	Name            string    `gorm:"unique_index;not null"`
	Description     string
	Image           string
	Visibility      string
	MetaTitle       string
	MetaDescription string
	CreateBy        uint
	UpdateBy        uint
}

// IsNode IsNode
func (v *Tag) IsNode() {}

// BeforeCreate 初始化uuid
func (v *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedTags PreDefinedTags
var PreDefinedTags = []*Tag{
	{Name: "article", Slug: "article", Description: "博客文章"},
	{Name: "learn", Slug: "learn", Description: "学习笔记"},
	{Name: "goweb", Slug: "goweb", Description: "动态网站"},
	{Name: "python", Slug: "python", Description: "python"},
	{Name: "javascript", Slug: "javascript", Description: "javascript"},
}
