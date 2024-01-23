package routes

import (
	"github.com/dzakimaulana/golaundry/internal/items"
	"github.com/gofiber/fiber/v2"
)

func ItemRouter(h *items.Handler, f *fiber.App) {
	item := f.Group("/api/item")
	item.Post("/add", h.AddItem)
	item.Get("/get", h.GetItem)
	item.Put("/update", h.UpdateItem)
	item.Delete("/delete", h.DeleteItem)
}
