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
	handle := handler.InitAuthHandler(service)

	authRouter := app.Group("api/v1/auth")

	adminGroup := app.Group("api/v1/auth/admin")
	adminGroup.Use(middleware.RoleAuthorizeMiddleware("Admin"))

	authRouter.Post("/signup", handle.SignUpUser)

	authRouter.Post("/login", handle.LogInUser)

	adminGroup.Get("/get", middleware.VerifyToken, handle.GetUser)
	adminGroup.Get("/get-id/:id", middleware.VerifyToken, handle.SelectUser)
	adminGroup.Patch("/update/:id", middleware.VerifyToken, handle.UpdateUser)
	adminGroup.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteUser)

}
