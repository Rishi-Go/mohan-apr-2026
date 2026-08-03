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

	likeRouter := app.Group("api/v1/like")

	userGroup := app.Group("api/v1/like/user")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User"), middleware.AuthUserMiddleware())

	adminGroup := app.Group("api/v1/like/admin")
	adminGroup.Use(middleware.AuthUserMiddleware())

	userGroup.Post("/insert", middleware.VerifyToken, handle.InsertLike) //user
	likeRouter.Get("/get", middleware.VerifyToken, handle.GetLike)
	likeRouter.Get("/get-id/:id", middleware.VerifyToken, handle.SelectLike)
	adminGroup.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteLike) //user
}
