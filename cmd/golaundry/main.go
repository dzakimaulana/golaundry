package main

import (
	"log"
	"time"

	"github.com/dzakimaulana/golaundry/internal/customers"
	"github.com/dzakimaulana/golaundry/internal/items"
	"github.com/dzakimaulana/golaundry/internal/transactions"
	"github.com/dzakimaulana/golaundry/internal/users"
	"github.com/dzakimaulana/golaundry/pkg/database"
	"github.com/dzakimaulana/golaundry/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	dbConn, err := database.DatabaseConn()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	store := session.New(session.Config{
		Expiration:   1 * time.Hour,
		KeyLookup:    "cookie:my-session",
		KeyGenerator: utils.UUIDv4,
		CookieName:   "my-session",
		CookieSecure: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Allow all origins
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           3600, // Maximum cache age in seconds
	}))

	userRepo := users.NewRepository(dbConn.GetDB())
	userSvc := users.NewService(userRepo)
	userHandler := users.NewHandler(userSvc, store)
	routes.UserRouter(userHandler, app, store)

	itemRepo := items.NewRepository(dbConn.GetDB())
	itemSvc := items.NewService(itemRepo)
	itemHandler := items.NewHandler(itemSvc)
	routes.ItemRouter(itemHandler, app, store)

	transactionsRepo := transactions.NewRepository(dbConn.GetDB())
	transactionsSvc := transactions.NewService(transactionsRepo)
	transactionsHandler := transactions.NewHandler(transactionsSvc)
	routes.TransactionsRouter(transactionsHandler, app, store)

	customersRepo := customers.NewRepository(dbConn.GetDB())
	customersSvc := customers.NewService(customersRepo)
	customersHandler := customers.NewHandler(customersSvc)
	routes.CustomersRouter(customersHandler, app, store)

	err = app.Listen(":3000")
	if err != nil {
		panic("Error starting server")
	}
}
