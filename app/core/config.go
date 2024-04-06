package core

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config is the struct for the configuration
type ConfYaml struct {
	AppEnv            string `mapstructure:"app_env"`
	Host              string `mapstructure:"host"`
	Port              string `mapstructure:"port"`
	DBDSN             string `mapstructure:"db_dsn"`
	AppID             string `mapstructure:"app_id"`
	AppSecret         string `mapstructure:"app_secret"`
	DefaultUserPrefix string `mapstructure:"default_user_prefix"`
	RedisAddr         string `mapstructure:"redis_addr"`
	RedisPassword     string `mapstructure:"redis_password"`
	RedisDB           int    `mapstructure:"redis_db"`
}

var conf *ConfYaml

func (c *ConfYaml) GetGinMode() string {
	switch c.AppEnv {
	case "dev":
		return gin.DebugMode
	case "development":
		return gin.DebugMode
	case "test":
		return gin.DebugMode
	case "prod":
		return gin.ReleaseMode
	case "production":
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}

func (c *ConfYaml) GetAddr() string {
	return c.Host + ":" + c.Port
}

func InitConf() *ConfYaml {

	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath("../config/")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	conf = &ConfYaml{}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	return conf
}

func GetConf() *ConfYaml {
	return conf
}
