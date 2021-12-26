package http

import (
	fiber "github.com/gofiber/fiber/v2"
	"go-otoklix-blog/infra/common"
)

// Routes ...
func Routes(app *fiber.App) {

	app.Get(common.APP_CTX+"/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	BlogPostRoutes(app, common.API_CTX)
}
