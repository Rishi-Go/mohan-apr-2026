package router

import (
	"blog_post/internals/handler"
	"blog_post/internals/middleware"
	"blog_post/internals/repository"
	"blog_post/internals/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetCategoryRouter(app fiber.Router, Db *gorm.DB) {

	repo := repository.InitCategoryRepo(Db)
	service := service.InitCategoryService(repo)
	handle := handler.InitCategoryHandler(service)

	categoryRouter := app.Group("api/v1/category")

	adminGroup := app.Group("api/v1/category/admin")
	adminGroup.Use(middleware.RoleAuthorizeMiddleware("Admin"))

	adminGroup.Post("/insert", middleware.VerifyToken, handle.InsertCategory) //only admin
	categoryRouter.Get("/get", middleware.VerifyToken, handle.GetCategory)
	categoryRouter.Get("/get-id/:id", middleware.VerifyToken, handle.SelectCategory)
	adminGroup.Patch("/update/:id", middleware.VerifyToken, handle.UpdateCategory)  //only admin
	adminGroup.Delete("/delete/:id", middleware.VerifyToken, handle.DeleteCategory) //only admin

}
