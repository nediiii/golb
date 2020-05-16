package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Setting 系统配置
type Setting struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;unique_index"`
	Key         string    `gorm:"unique_index;not null"`
	Value       string
	Description string
	CreateBy    uint
	UpdateBy    uint
}

// IsNode IsNode
func (v *Setting) IsNode() {}

// BeforeCreate 初始化uuid
func (v *Setting) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedSettings predefined settings
var PreDefinedSettings = []Setting{
	{Key: "systemInitTime", Value: time.Now().String(), Description: "first time system init"},
	{Key: "title", Value: "My Blog", Description: "blog global title"},
	{Key: "description", Value: "Yet another blog", Description: "blog global description"},
	{Key: "email", Value: "blog@example.com", Description: "blog administrator email"},
	{Key: "logo", Value: "/public/images/blog-logo.jpg", Description: "blog global logo"},
	{Key: "cover", Value: "/public/images/blog-cover.jpg"},
	{Key: "postPerPage", Value: "5"},
}
