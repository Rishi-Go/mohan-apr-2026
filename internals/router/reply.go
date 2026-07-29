package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetReplyRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitReplyRepo(Db)
	service := service.InitReplyService(repo)
	handle := handler.InitReplyHandler(service)

	replyRouter := app.Group("api/v1/reply")

	replyRouter.Post("/insert", handle.InsertReply)
	replyRouter.Get("/get", handle.GetReply)
	// replyRouter.Get("/get-id/:id", handle.)
	// replyRouter.Patch("/update/:id", handle.UpdateLike)
	// replyRouter.Delete("/delete/:id", handle.DeleteLike)
}
