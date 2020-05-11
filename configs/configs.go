package configs

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// DatabaseConfig DatabaseConfig
type DatabaseConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	Debug    bool
}

func init() {
	readConfigFile()
}

func readConfigFile() {
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
}

// ReadConfig ReadConfig
func (dc *DatabaseConfig) ReadConfig() {
	prefix := "database"
	viper.UnmarshalKey(prefix, dc)
}

// GetJwtKey GetJwtKey
func GetJwtKey() (key string) {
	viper.UnmarshalKey("site.jwt.key", key)
	return
}