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

	categoryRouter := app.Group("/category")
	categoryRouter.Use(middleware.VerifyToken)

	adminGroup := app.Group("/admin/category")
	adminGroup.Use(middleware.RoleAuthorizeMiddleware("Admin"), middleware.VerifyToken)
	
	//routes
	
	categoryRouter.Get("/get", handle.GetCategory)
	categoryRouter.Get("/get-id/:id", handle.SelectCategory)

	adminGroup.Post("/insert", handle.InsertCategory) //only admin
	adminGroup.Patch("/update/:id", handle.UpdateCategory)  //only admin
	adminGroup.Delete("/delete/:id", handle.DeleteCategory) //only admin

}
