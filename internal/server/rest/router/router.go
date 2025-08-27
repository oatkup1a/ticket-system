package router

import (
	"github.com/gofiber/fiber/v2"
	"oatkup.sys/internal/server/rest/handler"
)

func MountApplicationRoutes(app *fiber.App, h *handler.ApplicationHandler) {
	app.Post("/applications", h.Create)
}
