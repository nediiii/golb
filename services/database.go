package services

import (
	"fmt"

	"golb/configs"
	"golb/models"

	"log"

	"github.com/jinzhu/gorm"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// TODO close the connection

// DB 数据库连接实例
var DB *gorm.DB

func init() {
	establishConnection()
	InitDatabase()
}

// InitDatabase init database table and data
func InitDatabase() {
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
	db.LogMode(config.Debug)
	DB = db
}

// init database table
func initTable() {
	DB.AutoMigrate(&models.Setting{}, &models.User{}, &models.Role{}, &models.Post{}, &models.Tag{}, &models.Comment{})
}

// init database data
func initData() {
	// init system settings data
	for _, s := range models.PreDefinedSettings {
		err := DB.Unscoped().Where(models.Setting{Key: s.Key}).FirstOrCreate(&s).GetErrors()
		for _, e := range err {
			log.Fatal(e.Error())
		}
	}

	// init roles data
	for _, r := range models.PreDefinedRoles {
		err := DB.Unscoped().Where(models.Role{Name: r.Name}).FirstOrCreate(&r).GetErrors()
		for _, e := range err {
			log.Fatal(e.Error())
		}
	}

	// init users data
	for _, u := range models.PreDefinedUsers {
		err := DB.Unscoped().Where(models.User{Slug: u.Slug}).FirstOrCreate(&u).GetErrors()
		for _, e := range err {
			log.Fatal(e.Error())
		}
	}

	// init tags data
	for _, t := range models.PreDefinedTags {
		err := DB.Unscoped().Where(models.Tag{Name: t.Name}).FirstOrCreate(&t).GetErrors()
		for _, e := range err {
			log.Fatal(e.Error())
		}
	}

	// init posts data
	for _, p := range models.PreDefinedPosts {
		err := DB.Set("gorm:association_autoupdate", false).Unscoped().Where(models.Post{Slug: p.Slug}).FirstOrCreate(&p).GetErrors()
		for _, e := range err {
			log.Fatal(e.Error())
		}
	}
}
