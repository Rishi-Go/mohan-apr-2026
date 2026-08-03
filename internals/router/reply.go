package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
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

	userGroup := app.Group("api/v1/reply/user")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User"), middleware.AuthUserMiddleware())

	adminGroup := app.Group("api/v1/reply/admin")
	adminGroup.Use(middleware.AuthUserMiddleware())

	userGroup.Post("/insert", handle.InsertReply) //Author
	replyRouter.Get("/get", handle.GetReply)
	replyRouter.Get("/get-id/:id", handle.SelectReply)
	userGroup.Patch("/update/:id", handle.UpdateReply) //Author
	adminGroup.Delete("/delete/:id", handle.DeleteReply) //Author & Admin
}
