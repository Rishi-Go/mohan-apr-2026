package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetLikeRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitLikeRepo(Db)
	service := service.InitLikeService(repo)
	handle := handler.InitLikeHandler(service)

	likeRouter := app.Group("/like")
	likeRouter.Use(middleware.VerifyToken)

	userGroup := app.Group("/user/like")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User","Admin"), middleware.AuthUserMiddleware(), middleware.VerifyToken)


	userGroup.Post("/insert/:id", handle.InsertLike) //user
	likeRouter.Get("/get", handle.GetLike)
	likeRouter.Get("/get-id/:id", handle.SelectLike)
	userGroup.Delete("/delete/:id", handle.DeleteLike) //user
}
