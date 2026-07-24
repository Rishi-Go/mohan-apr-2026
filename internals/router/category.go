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

	authRouter := app.Group("api/v1/category")

	authRouter.Post("/insert", handle.InsertCategory)
	authRouter.Get("/get",handle.GetCategory)
	authRouter.Get("/get-id/:id",handle.SelectCategory)
	authRouter.Patch("/update/:id", handle.UpdateCategory)
	authRouter.Delete("/delete/:id",handle.DeleteCategory)
	
}