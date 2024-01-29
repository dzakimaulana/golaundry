package routes

import (
	"github.com/dzakimaulana/golaundry/internal/users"
	"github.com/dzakimaulana/golaundry/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func UserRouter(h *users.Handler, f *fiber.App, session *session.Store) {
	user := f.Group("/api/user")
	user.Post("/login", h.Login)
	user.Post("/create-user", h.CreateUser)
	user.Get("/get", middlewares.GetSession(session), middlewares.OnlyAdmin, h.GetAllUser)
	user.Get("/get/:id", middlewares.GetSession(session), middlewares.OnlyAdmin, h.GetUserByID)
	user.Put("/reset-password", middlewares.GetSession(session), h.ResetPassword)
	user.Post("/logout", middlewares.GetSession(session), h.Logout)
}
