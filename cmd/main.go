package main

import (
	"blog_post/config"
	"blog_post/drivers/db"
	"fmt"
)

func main() {
	cfg := config.InitConfig()

	_, err := db.InitDB(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// router := router.SetRouter(Db)

	// fmt.Println("Server started ...")

	// log.Fatal(http.ListenAndServe(cfg.HttpPortConfig.HttpPort, router))
}
