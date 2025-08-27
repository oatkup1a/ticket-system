package rest

import (
	"github.com/gofiber/fiber/v2"
	"oatkup.sys/internal/server/rest/handler"
	"oatkup.sys/internal/server/rest/middleware"
	"oatkup.sys/internal/server/rest/router"
)

func NewServer(appHandler *handler.ApplicationHandler) *fiber.App {
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recovery())

	// app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Static("/swagger", "./swagger")
	app.Static("/openapi.yaml", "./docs/openapi.yaml")

	router.MountApplicationRoutes(app, appHandler)
	return app
}
