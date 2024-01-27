package routes

import (
	"github.com/dzakimaulana/golaundry/internal/transactions"
	"github.com/gofiber/fiber/v2"
)

func TransactionRouter(h *transactions.Handler, f *fiber.App) {
	transaction := f.Group("/api/transaction")
	transaction.Post("/add", h.AddTs)
	transaction.Get("/get", h.GetAllTs)
	// transaction.Put("/update")
	// transaction.Delete("/delete")
}
