package services

import (
	"fmt"

	"golb/configs"
	"golb/models"

	"github.com/jinzhu/gorm"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TODO close the connection

// DB 数据库连接实例
var DB *gorm.DB

func init() {
	establishConnection()
	initTable()
	initData()
}

// connect to database
func establishConnection() {
	config := &configs.DatabaseConfig{}
	config.ReadConfig()
	connectionURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.User, config.DBName, config.Password, "disable")
	// fmt.Println(connectionURL)
	db, err := gorm.Open("postgres", connectionURL)
	if err != nil {
		println(err.Error())
		panic("failed to connect database")
	}
	db.Set("gorm:association_autoupdate", false)
	db.LogMode(config.Debug)
	DB = db
}

// init database table
func initTable() {
	DB.AutoMigrate(&models.Setting{}, &models.User{}, &models.Role{}, &models.Post{}, &models.Tag{})
}

// init database data
func initData() {
	// init system settings data
	for _, s := range models.PreDefinedSettings {
		DB.Where(models.Setting{Key: s.Key}).FirstOrCreate(&s)
	}

	// init roles data
	for _, r := range models.PreDefinedRoles {
		DB.Where(models.Role{Name: r.Name}).FirstOrCreate(&r)
	}

	// init users data
	for _, u := range models.PreDefinedUsers {
		DB.Where(models.User{Name: u.Name}).FirstOrCreate(&u)
	}

	// init tags data
	for _, t := range models.PreDefinedTags {
		DB.Where(models.Tag{Name: t.Name}).FirstOrCreate(&t)
	}

	// init posts data
	for _, p := range models.PreDefinedPosts {
		DB.Where(models.Post{Slug: p.Slug}).FirstOrCreate(&p)
	}
}
