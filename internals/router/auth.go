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

	authRouter := app.Group("/auth")
	authRouter.Use(middleware.VerifyToken)

	userGroup := app.Group("/user/auth")
	userGroup.Use(middleware.VerifyToken, middleware.AuthUserMiddleware())

	adminGroup := app.Group("/admin/auth")
	adminGroup.Use(middleware.VerifyToken, middleware.RoleAuthorizeMiddleware("Admin"))

	// routes

	authRouter.Post("/signup", handle.SignUpUser)
	authRouter.Post("/login", handle.LogInUser)

	adminGroup.Get("/get", handle.GetUser)
	adminGroup.Get("/get-id/:id", handle.SelectUser)
	
	userGroup.Patch("/update/:id", handle.UpdateUser)
	userGroup.Delete("/delete/:id", handle.DeleteUser)

}
