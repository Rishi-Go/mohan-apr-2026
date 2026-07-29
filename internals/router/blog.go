package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetBlogRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitBlogRepo(Db)
	service := service.InitBlogService(repo)
	handle := handler.InitBlogHandler(service)

	blogRouter := app.Group("api/v1/blog")

	blogRouter.Post("/insert", middleware.VerifyToken, handle.InsertBlog)
	blogRouter.Get("/get",middleware.VerifyToken, handle.GetBlog)
	blogRouter.Get("/get-id/:id",middleware.VerifyToken, handle.SelectBlog)
	blogRouter.Patch("/update/:id",middleware.VerifyToken, handle.UpdateBlog)
	blogRouter.Delete("/delete/:id",middleware.VerifyToken, handle.DeleteBlog)
}
