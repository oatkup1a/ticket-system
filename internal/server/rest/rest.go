package rest

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"oatkup.sys/internal/server/rest/handler"
	"oatkup.sys/internal/server/rest/middleware"
	"oatkup.sys/internal/server/rest/router"
)

func NewServer(appHandler *handler.ApplicationHandler) *fiber.App {
	app := fiber.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recovery())

	// Debug: print working directory and check files
	wd, _ := os.Getwd()
	log.Printf("Working directory: %s", wd)
	
	swaggerPath := filepath.Join(wd, "swagger")
	openapiPath := filepath.Join(wd, "docs", "openapi.yaml")
	
	if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
		log.Printf("WARNING: swagger directory not found at: %s", swaggerPath)
	} else {
		log.Printf("✓ Swagger directory found at: %s", swaggerPath)
	}
	
	if _, err := os.Stat(openapiPath); os.IsNotExist(err) {
		log.Printf("WARNING: openapi.yaml not found at: %s", openapiPath)
	} else {
		log.Printf("✓ OpenAPI spec found at: %s", openapiPath)
	}

	// Serve Swagger UI and OpenAPI spec
	app.Static("/swagger", "./swagger")
	app.Get("/openapi.yaml", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/openapi.yaml")
	})

	router.MountApplicationRoutes(app, appHandler)
	return app
}