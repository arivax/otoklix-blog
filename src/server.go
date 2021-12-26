package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"go-otoklix-blog/infra/common"
	"go-otoklix-blog/infra/config"

	"go-otoklix-blog/delivery/http"
)

func main() {
	common.ShowBanner()

	// Setup http server
	app := fiber.New()

	http.Routes(app)

	// listen/Serve the new Fiber app on port HTTP_PORT
	log.Fatal(app.Listen(":" + config.Get("HTTP_PORT")))
}
