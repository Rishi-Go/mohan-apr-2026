package main

import (
	"blog_post/config"
	"blog_post/drivers/db"
	"blog_post/internals/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
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
