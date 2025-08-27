package main

import (
	"log"
	"os"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	"oatkup.sys/internal/di"
	"oatkup.sys/internal/server/rest"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN") // e.g. postgres DSN
	if dsn == "" {
		log.Fatal("DATABASE_DSN is empty")
	}

	deps, err := di.NewAppDeps(dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := rest.NewServer(deps.AppHandler)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
