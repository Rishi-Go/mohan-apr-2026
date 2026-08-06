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

	replyRouter := app.Group("/reply")
	replyRouter.Use(middleware.VerifyToken)

	userGroup := app.Group("/user/reply")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User","Admin"), middleware.AuthUserMiddleware(), middleware.VerifyToken)



	userGroup.Post("/insert/:id", handle.InsertReply) //comment(owner)
	replyRouter.Get("/get", handle.GetReply)
	replyRouter.Get("/get-id/:id", handle.SelectReply)
	userGroup.Patch("/update/:id", handle.UpdateReply) //Author
	userGroup.Delete("/delete/:id", handle.DeleteReply) //Author & Admin
}
