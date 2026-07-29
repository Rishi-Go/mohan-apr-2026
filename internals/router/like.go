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

	likeRouter.Post("/insert", middleware.VerifyToken, handle.InsertLike)
	likeRouter.Get("/get", middleware.VerifyToken, handle.GetLike)
	likeRouter.Get("/get-id/:id", middleware.VerifyToken, handle.SelectLike)
	likeRouter.Patch("/update/:id", middleware.VerifyToken, handle.UpdateLike)
	likeRouter.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteLike)
}
