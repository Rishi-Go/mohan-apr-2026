package main

import (
	"blog_post/config"
	"blog_post/drivers/db"
	"blog_post/internals/router"
	"blog_post/pkg/logger"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {

	//initializing Zaplogger
	logger.Init("development")
	defer logger.Sync()

	// Call structured log entries using strictly typed fields
	logger.Log.Info("Server starting up",
		zap.String("port", ":8080"),
	)

	// initialize configs
	cfg := config.InitConfig()

	// initialize DB
	Db, err := db.InitDB(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()

	router.SetRouter(app, Db)

	fmt.Println("Server started ...")

	log.Fatal(app.Listen(cfg.HttpPortConfig.HttpPort))
}
