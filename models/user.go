package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User user model
type User struct {
	gorm.Model
	Roles           []*Role   `gorm:"many2many:user_roles"`
	Posts           []*Post   `gorm:"many2many:user_posts"`
	UUID            uuid.UUID `gorm:"type:uuid;unique_index"`
	Slug            string    `gorm:"unique_index;not null"`
	Name            string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	Email           string
	Image           string
	Cover           string
	Bio             string
	Website         string
	Location        string
	Accessibility   string
	Status          string
	Language        string
	Visibility      string
	MetaTitle       string
	MetaDescription string
	LastLogin       time.Time
	CreateBy        uint
	UpdateBy        uint
}

// GetID GetID
func (v User) GetID() interface{} {
	return v.ID
}

// GetCreateAt GetCreateAt
func (v User) GetCreateAt() interface{} {
	return v.CreatedAt
}

// GetUpdateAt GetUpdateAt
func (v User) GetUpdateAt() interface{} {
	return v.UpdatedAt
}

// BeforeCreate 初始化uuid
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

// PreDefinedUsers PreDefinedUsers
var PreDefinedUsers = []*User{
	{Name: "owner", Password: "root", Roles: PreDefinedRoles},
	{Name: "admin", Password: "admin", Roles: PreDefinedRoles[1:]},
}
