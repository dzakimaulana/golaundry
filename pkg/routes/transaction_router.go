package routes

import (
	"github.com/dzakimaulana/golaundry/internal/transactions"
	"github.com/dzakimaulana/golaundry/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func TransactionsRouter(h *transactions.Handler, f *fiber.App, session *session.Store) {
	transaction := f.Group("/api/transaction")
	transaction.Use(middlewares.GetSession(session))
	transaction.Post("/add", h.AddTs)
	transaction.Get("/get", h.GetAllTs)
	transaction.Get("/get/:id", h.GetTsByID)
}
