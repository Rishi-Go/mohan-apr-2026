package db

import (
	"blog_post/config"
	"blog_post/pkg/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {

	dst := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", cfg.DataBaseConfig.Host, cfg.DataBaseConfig.User, cfg.DataBaseConfig.Password, cfg.DataBaseConfig.DbName, cfg.DataBaseConfig.Port, cfg.DataBaseConfig.SslMode)

	db, err := gorm.Open(postgres.Open(dst), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database:%w", err)
	}

	err = db.AutoMigrate(&models.BlogUsers{},&models.Blog{},&models.Category{},&models.Comment{},&models.Like{},&models.Reply{})
	if err != nil {
		return nil, fmt.Errorf("Migration Error: %v", err)
	}
	return db, nil

}
