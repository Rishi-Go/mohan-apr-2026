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

	commentRouter.Post("/insert", middleware.VerifyToken, handle.InsertComment)
	commentRouter.Get("/get", middleware.VerifyToken, handle.GetComment)
	commentRouter.Get("/get-id/:id", middleware.VerifyToken, handle.SelectComment)
	commentRouter.Patch("/update/:id", middleware.VerifyToken, handle.UpdateComment)
	commentRouter.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteComment)

}
