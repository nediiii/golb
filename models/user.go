package models

import (
	"errors"
	"log"
	"time"

	"golb/utils"

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

// IsNode IsNode
func (v *User) IsNode() {}

// BeforeCreate 初始化uuid
func (v *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("UUID", uuid.New())
}

// BeforeSave 初始化uuid
func (v *User) BeforeSave(scope *gorm.Scope) error {
	if pwdLen := len(v.Password); pwdLen >= 8 {
		scope.SetColumn("Password", utils.Hash(v.Password))
		log.Println("user save ", utils.Hash(v.Password))
		return nil
	}
	return errors.New("password length should longer than 8 characters")
}

// PreDefinedUsers PreDefinedUsers
var PreDefinedUsers = []*User{
	{Name: "owner", Slug: "owner", Password: "rootroot", Roles: PreDefinedRoles},
	{Name: "admin", Slug: "admin", Password: "adminadmin", Roles: PreDefinedRoles[1:]},
}
