package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func SetAuthRouter(app fiber.Router, Db *gorm.DB) {
	authRouter := app.Group("/auth")

	repo := repository.InitAuthRepo(Db)
	service := service.InitAuthService(repo)
	handle := handler.InitHandler(service)

	authRouter.Post("/insert", handle.InsertSignUP)
}
