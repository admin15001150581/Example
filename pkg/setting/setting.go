package setting

import (
	"github.com/spf13/viper"
	"log"
)

var Viper *viper.Viper
func init()  {
	Viper =viper.New()
	Viper.SetConfigName("config")
	Viper.SetConfigType("toml")
	Viper.AddConfigPath("config/")
	Viper.SetDefault("redis.port",6379)
	errConfig:= Viper.ReadInConfig()
	if errConfig !=nil{
		log.Fatal("read config failed %v",errConfig)
	}
}