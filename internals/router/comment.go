package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetCommentRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitCommentRepo(Db)
	service := service.InitCommentService(repo)
	handle := handler.InitCommentHandler(service)

	commentRouter := app.Group("api/v1/comment")

	userGroup := app.Group("api/v1/comment/user")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User"), middleware.AuthUserMiddleware())

	adminGroup := app.Group("api/v1/comment/admin")
	adminGroup.Use(middleware.AuthUserMiddleware())

	userGroup.Post("/insert", middleware.VerifyToken, handle.InsertComment) //(only user)
	commentRouter.Get("/get", middleware.VerifyToken, handle.GetComment)
	commentRouter.Get("/get-id/:id", middleware.VerifyToken, handle.SelectComment)
	userGroup.Patch("/update/:id", middleware.VerifyToken, handle.UpdateComment)   // (only user)
	adminGroup.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteComment) // Admin & User
}
