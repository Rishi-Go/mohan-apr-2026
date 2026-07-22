package router

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

func SetRouter(app *fiber.App, Db *gorm.DB) {
	apiRouter := app.Group("/api/v1")
	SetAuthRouter(apiRouter, Db)
}
