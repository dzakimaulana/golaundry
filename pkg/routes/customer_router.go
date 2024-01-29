package routes

import (
	"github.com/dzakimaulana/golaundry/internal/customers"
	"github.com/dzakimaulana/golaundry/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CustomersRouter(h *customers.Handler, f *fiber.App, session *session.Store) {
	customer := f.Group("/api/customer")
	customer.Use(middlewares.GetSession(session))
	customer.Post("/add", middlewares.OnlyAdmin, h.AddNewCustomer)
	customer.Get("/get", h.GetAllCustomer)
	customer.Get("/get/:id", h.GetCustomerByID)
}
