package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var App Config

type Config struct {
	DatabaseType     string `mapstructure:"DB_TYPE"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     int    `mapstructure:"DB_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`

	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
	S3Region     string `mapstructure:"S3_REGION"`
	S3APIKey     string `mapstructure:"S3_API_KEY"`
	S3SecretKey  string `mapstructure:"S3_SECRET_KEY"`
	S3Domain     string `mapstructure:"S3_DOMAIN"`

	JwtSecret string `mapstructure:"JWT_SECRET"`

	Port string `mapstructure:"PORT"`
}

func init() {
	LoadedConfAppFromTheEnv, err := loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config.\n", err)
	}

	App = LoadedConfAppFromTheEnv

	fmt.Println("loaded config!")
}

func loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("cannot load .env file.\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("unable to decode into struct.\n", err)
	}

	return
}
