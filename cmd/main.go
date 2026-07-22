package main

import (
	"blog_post/config"
	"blog_post/drivers/db"
	"blog_post/internals/router"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	cfg := config.InitConfig()
	Db, err := db.InitDB(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()

	router.SetAuthRouter(app, Db)

	fmt.Println("Server started ...")

	log.Fatal(app.Listen(cfg.HttpPortConfig.HttpPort))
}
