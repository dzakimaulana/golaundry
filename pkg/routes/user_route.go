package routes

import (
	"github.com/dzakimaulana/golaundry/internal/users"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(h *users.Handler, f *fiber.App) {
	user := f.Group("/api/user")
	user.Get("/login")
	user.Get("/create-user", h.CreateUser)
}
