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

	commentRouter := app.Group("comment")
	commentRouter.Use(middleware.VerifyToken)

	userGroup := app.Group("user/comment")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User", "Admin"), middleware.AuthUserMiddleware(), middleware.VerifyToken)


	userGroup.Post("/insert", handle.InsertComment) //(only user)
	commentRouter.Get("/get", handle.GetComment)
	commentRouter.Get("/get-id/:id", handle.SelectComment)
	userGroup.Patch("/update/:id", handle.UpdateComment)   // (only user)
	userGroup.Delete("/delete/:id",  handle.DeleteComment) // Admin & User
}
