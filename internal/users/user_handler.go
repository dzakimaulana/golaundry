package users

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserService
}

func NewHandler(s UserService) *Handler {
	return &Handler{
		UserService: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var u CreateUserReq
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	res, err := h.UserService.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var u LoginReq
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	res, err := h.UserService.Login(c.Context(), &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
