package routes

import (
	"github.com/dzakimaulana/golaundry/internal/items"
	"github.com/dzakimaulana/golaundry/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func ItemRouter(h *items.Handler, f *fiber.App, session *session.Store) {
	item := f.Group("/api/item")
	item.Use(middlewares.GetSession(session))
	item.Post("/add", middlewares.OnlyAdmin, h.AddItem)
	item.Get("/get", h.GetItem)
	item.Get("/get/:id", h.GetItemByID)
	item.Put("/update", middlewares.OnlyAdmin, h.UpdateItem)
	item.Delete("/delete/:id", middlewares.OnlyAdmin, h.DeleteItem)
}
