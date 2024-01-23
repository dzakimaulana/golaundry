package main

import (
	"log"

	"github.com/dzakimaulana/golaundry/internal/items"
	"github.com/dzakimaulana/golaundry/internal/users"
	"github.com/dzakimaulana/golaundry/pkg/database"
	"github.com/dzakimaulana/golaundry/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	dbConn, err := database.DatabaseConn()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	app := fiber.New()

	userRepo := users.NewRepository(dbConn.GetDB())
	userSvc := users.NewService(userRepo)
	userHandler := users.NewHandler(userSvc)
	routes.UserRouter(userHandler, app)

	itemRepo := items.NewRepository(dbConn.GetDB())
	itemSvc := items.NewService(itemRepo)
	itemHandler := items.NewHandler(itemSvc)
	routes.ItemRouter(itemHandler, app)

	err = app.Listen(":3000")
	if err != nil {
		panic("Error starting server")
	}
}
