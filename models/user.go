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
	Posts           []*Post   `gorm:"many2many:user_posts"`
	Role            *Role     `gorm:"foreignkey:RoleID"` // Belongs to: `User` Belong to one `Role` , The ForeignKey define in `User`
	UUID            uuid.UUID `gorm:"type:uuid;unique_index"`
	Slug            string    `gorm:"unique_index;not null"`
	Name            string    `gorm:"not null"`
	Password        string    `gorm:"not null"`
	RoleID          uint
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
	// 经过bcrypt的密码长度为120 -> 不再进行bcrypt,不需要操作
	// 允许的密码长度为8到64 -> 进行bcrypt后存入
	// 不允许其他长度的密码 -> 拒绝操作,抛出错误
	log.Println("BeforeSave trigger")
	passwordLen := len(v.Password)
	if passwordLen >= 8 && passwordLen <= 64 {
		log.Println("BeforeSave trigger:", "operate to hash the password")
		scope.SetColumn("Password", utils.Hash(v.Password))
	} else if passwordLen == 120 {
		log.Println("BeforeSave trigger:", "no operate need")
		// nothing to do
	} else {
		log.Println("BeforeSave trigger:", "password length wrong")
		return errors.New("password length invalid")
	}
	return nil
}

// PreDefinedUsers PreDefinedUsers
var PreDefinedUsers = []*User{
	{Name: "owner", Slug: "owner", Password: "adminadmin", RoleID: 1},
	{Name: "admin", Slug: "admin", Password: "adminadmin", Role: PreDefinedRoles[1]},
	{Name: "editor", Slug: "admin", Password: "adminadmin", Role: PreDefinedRoles[2]},
	{Name: "author", Slug: "admin", Password: "adminadmin", RoleID: 4},
	{Name: "contributer", Slug: "admin", Password: "adminadmin", Role: PreDefinedRoles[4]},
}
