package configs

import (
	"log"
	"os"

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
	// make sure when testing it can find the config file
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

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
			log.Fatal(err, "config file not exist!(if you are running test, make sure you run it in project base path,like `golb`)")
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

// JwtKey JwtKey
func JwtKey() (key []byte) {
	var s string
	viper.UnmarshalKey("site.secure.jwt.key", &s)
	return []byte(s)
}

// JwtEnable JwtEnable
func JwtEnable() (enable bool) {
	viper.UnmarshalKey("site.secure.jwt.enable", &enable)
	return
}

// Pepper Pepper
func Pepper() (pepper string) {
	viper.UnmarshalKey("site.secure.pepper", &pepper)
	return
}
