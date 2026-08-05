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

	blogRouter := app.Group("blog")
	blogRouter.Use(middleware.VerifyToken)

	userGroup := app.Group("user/blog")
	userGroup.Use(middleware.RoleAuthorizeMiddleware("User", "Admin"), middleware.AuthUserMiddleware(), middleware.VerifyToken)

	userGroup.Post("/insert", handle.InsertBlog) // only by author (specific user who create the blog)
	blogRouter.Get("/get", handle.GetBlog)
	blogRouter.Get("/get-id/:id", handle.SelectBlog)
	userGroup.Patch("/update/:id", handle.UpdateBlog)  // only by author (specific user who create the blog)
	userGroup.Delete("/delete/:id", handle.DeleteBlog) // Admin & User
}
