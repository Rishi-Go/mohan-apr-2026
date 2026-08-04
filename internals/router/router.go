package router

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetRouter(app fiber.Router, Db *gorm.DB) {

	api := app.Group("api/v1")

	SetAuthRouter(api,Db)
	SetCategoryRouter(api,Db)
	SetLikeRouter(app,Db)
	SetBlogRouter(app,Db)
	SetCommentRouter(app,Db)
	SetReplyRouter(app,Db)

}
