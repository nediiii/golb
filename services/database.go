package services

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golb/models"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type databaseConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	Debug    bool
}

var config databaseConfig

// TODO close the connection

// DB 数据库连接实例
var DB *gorm.DB

func init() {
	getConfig()
	establishConnection()
	initTable()
	initData()
}

// 读取配置文件
func getConfig() {
	// determine whether run under test
	if strings.HasSuffix(os.Args[0], ".test") {
		viper.AddConfigPath("../")
	} else {
		viper.AddConfigPath("./")
	}

	viper.SetConfigType("yaml")

	configName := "config.production"
	// 如果不是生产环境, 则加载本地的开发配置`config.development.yml`
	if os.Getenv("GIN_MODE") != "release" {
		configName = "config.development"
	}
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal(err, "config file not exist!")
		} else {
			// Config file was found but another error was produced
			log.Fatal(err, "could not read config file")
		}
	}
	prefix := "database"
	viper.UnmarshalKey(prefix, &config)
}

// connect to database
func establishConnection() {
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
