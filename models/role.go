package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Roles 和权限
// Contributors: Can log in and write posts, but cannot publish
// Authors: Can create and publish new posts and tags
// Editors: Can invite, manage and edit authors and contributors
// Administrators: Have full permissions to edit all data and settings
// Owner: An admin who cannot be deleted and has access to billing details

// Role 角色
type Role struct {
	gorm.Model  `gorm:"embedded"`
	Users       []*User   `gorm:"many2many:user_roles"`
	UUID        uuid.UUID `gorm:"type:uuid;unique_index"`
	Name        string    `gorm:"unique_index;not null"`
	Description string
	CreateBy    uint
	UpdateBy    uint
}

// IsNode IsNode
func (v *Role) IsNode() {}

// BeforeCreate 初始化uuid
func (v *Role) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedRoles 预定义的Roles
var PreDefinedRoles = []*Role{
	{Name: "Owner", Description: "An admin who cannot be deleted and has access to billing details"},
	{Name: "Administrator", Description: "Have full permissions to edit all data and settings"},
	{Name: "Editor", Description: "Can invite, manage and edit authors and contributors"},
	{Name: "Author", Description: "Can create and publish new posts and tags"},
	{Name: "Contributor", Description: "Can log in and write posts, but cannot publish"},
}
