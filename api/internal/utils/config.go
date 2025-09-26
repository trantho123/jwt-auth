package utils

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

type Config struct {
	PORT                string `mapstructure:"PORT"`
	MONGO_DATABSAE      string `mapstructure:"MONGO_DATABSAE"`
	MONGO_SERVER        string `mapstructure:"MONGO_SERVER"`
	MONGO_COLLECTION    string `mapstructure:"MONGO_COLLECTION"`
	REDIS_SERVER        string `mapstructure:"REDIS_SERVER"`
	EMAIL_FROM          string `mapstructure:"EMAIL_FROM"`
	SMTP_HOST           string `mapstructure:"SMTP_HOST"`
	SMTP_USER           string `mapstructure:"SMTP_USER"`
	SMTP_PASS           string `mapstructure:"SMTP_PASS"`
	CLIENT_ORIGIN       string `mapstructure:"CLIENT_ORIGIN"`
	PRIVATE_REFRESH_KEY string `mapstructure:"PRIVATE_REFRESH_KEY"`
	PUBLIC_REFRESH_KEY  string `mapstructure:"PUBLIC_REFRESH_KEY"`
	PRIVATE_ACCESS_KEY  string `mapstructure:"PRIVATE_ACCESS_KEY"`
	PUBLIC_ACCESS_KEY   string `mapstructure:"PUBLIC_ACCESS_KEY"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
