package router

import (
	"blog_post/internals/handler"
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

	categoryRouter.Post("/insert", handle.InsertCategory)
	categoryRouter.Get("/get",handle.GetCategory)
	categoryRouter.Get("/get-id/:id",handle.SelectCategory)
	categoryRouter.Patch("/update/:id", handle.UpdateCategory)
	categoryRouter.Delete("/delete/:id",handle.DeleteCategory)
	
}