package router

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetRouter(app fiber.Router, Db *gorm.DB) {

	SetAuthRouter(app,Db)
	SetCategoryRouter(app,Db)
	SetLikeRouter(app,Db)
	SetBlogRouter(app,Db)
	SetCommentRouter(app,Db)
	SetReplyRouter(app,Db)
}
