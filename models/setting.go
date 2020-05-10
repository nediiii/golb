package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Setting 系统配置
type Setting struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"type:uuid;unique_index"`
	Key      string    `gorm:"unique_index;not null"`
	Value    string
	Type     string
	CreateBy uint
	UpdateBy uint
}

// IsNode IsNode
func (v *Setting) IsNode() {}

// GetID GetID
func (v Setting) GetID() interface{} {
	return v.ID
}

// GetCreateAt GetCreateAt
func (v Setting) GetCreateAt() interface{} {
	return v.CreatedAt
}

// GetUpdateAt GetUpdateAt
func (v Setting) GetUpdateAt() interface{} {
	return v.UpdatedAt
}

// BeforeCreate 初始化uuid
func (setting *Setting) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedSettings 预设设置值
var PreDefinedSettings = []Setting{
	{UUID: uuid.New(), Key: "title", Value: "My Blog"},
	{UUID: uuid.New(), Key: "description", Value: "Yet another blog"},
	{UUID: uuid.New(), Key: "email", Value: "blog@example.com"},
	{UUID: uuid.New(), Key: "logo", Value: "/public/images/blog-logo.jpg"},
	{UUID: uuid.New(), Key: "cover", Value: "/public/images/blog-cover.jpg"},
	{UUID: uuid.New(), Key: "postPerPage", Value: "5"},
	{UUID: uuid.New(), Key: "activeTheme", Value: "default"},
	{UUID: uuid.New(), Key: "navigation", Value: `[{"label":"Home", "url":"/"}]`},
}
