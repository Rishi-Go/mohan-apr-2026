package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetAuthRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitAuthRepo(Db)
	service := service.InitAuthService(repo)
	handle:= handler.InitAuthHandler(service)

	authRouter := app.Group("api/v1/auth")

	authRouter.Post("/signup", handle.InsertUser)
	authRouter.Get("/get", handle.GetUser)
	authRouter.Get("/get-id/:id", handle.SelectUser)
	authRouter.Patch("/update/:id", handle.UpdateUser)
	authRouter.Delete("/delete/:id", handle.DeleteUser)

	authRouter.Post("/login",handle.LogInUser)
	authRouter.Get("/validate",middleware.RequestToken,handle.Validate)
}
//  api/v1/auth/login