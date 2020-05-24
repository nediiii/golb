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
	gorm.Model
	Users       []*User   `gorm:"foreignkey:RoleID"`
	UUID        uuid.UUID `gorm:"type:uuid;unique_index"`
	Name        string    `gorm:"unique_index;not null"`
	Description string
	CreateBy    uint
	UpdateBy    uint
}

// Role constant
const (
	RoleSys           = iota
	RoleOwner         // full permissions,only one user. can create other administrators
	RoleAdministrator // full permissions except operation that create owner role user or admin role user
	RoleEditor
	RoleAuthor      //
	RoleContributor // only have post permission, but can not publish a post
)

// IsNode IsNode
func (v *Role) IsNode() {}

// BeforeCreate 初始化uuid
func (v *Role) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedRoles should not modify
var PreDefinedRoles = []*Role{
	{Name: "Owner", Description: "An admin who cannot be deleted and has access to billing details"},
	{Name: "Administrator", Description: "Have full permissions to edit all data and settings"},
	{Name: "Editor", Description: "Can invite, manage and edit authors and contributors"},
	{Name: "Author", Description: "Can create and publish new posts and tags"},
	{Name: "Contributor", Description: "Can log in and write posts, but cannot publish"},
}
