package main

import (
	"log"
	"os"

	"oatkup.sys/internal/di"
	"oatkup.sys/internal/server/rest"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is empty")
	}

	deps, err := di.NewAppDeps(dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := rest.NewServer(deps.AppHandler)

	log.Println("Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}