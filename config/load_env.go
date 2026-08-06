package config

import (
	"blog_post/pkg/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DataBaseConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
}

type HttpPortConfig struct {
	HttpPort string
}

type Config struct {
	DataBaseConfig DataBaseConfig
	HttpPortConfig HttpPortConfig
}

func InitConfig() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		logger.Log.Warn("Error Loading .env file")
		log.Fatal("Error Loading .env file")
	}

	return &Config{
		DataBaseConfig: DataBaseConfig{
			Host:     GetEnv("DB_HOST", ""),
			User:     GetEnv("DB_USER", ""),
			Password: GetEnv("DB_PASSWORD", ""),
			DbName:   GetEnv("DB_NAME", ""),
			Port:     GetEnv("DB_PORT", ""),
			SslMode:  GetEnv("SSL_MODE", ""),
		},
		HttpPortConfig: HttpPortConfig{
			HttpPort: GetEnv("HTTP_PORT", ":8080"),
		},
	}

}

func GetEnv(str, defaultStr string) string {
	if str = os.Getenv(str); str != "" {
		return str
	}
	return defaultStr
}
