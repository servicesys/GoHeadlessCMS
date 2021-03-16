package storage

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func LoadConfig() Config {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config Config

	config = Config{
		Host:     viper.GetString("DB.HOST"),
		Port:     viper.GetInt("DB.PORT"),
		User:     viper.GetString("DB.USER"),
		Password: viper.GetString("DB.PASS"),
		DB:       viper.GetString("DB.NAME"),
	}
	return config
}
