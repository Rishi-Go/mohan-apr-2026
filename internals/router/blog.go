package router

import (
	"blog_post/internals/handler"
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

	blogRouter.Post("/insert", handle.InsertBlog)
	blogRouter.Get("/get", handle.GetBlog)
	blogRouter.Get("/get-id/:id", handle.SelectBlog)
	blogRouter.Patch("/update/:id", handle.UpdateBlog)
	blogRouter.Delete("/delete/:id", handle.DeleteBlog)
}
