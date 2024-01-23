package items

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	ItemService
}

func NewHandler(s ItemService) *Handler {
	return &Handler{
		ItemService: s,
	}
}

func (h *Handler) AddItem(c *fiber.Ctx) error {
	var item AddReq
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}
	res, err := h.ItemService.AddItem(c.Context(), &item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetItem(c *fiber.Ctx) error {
	name := c.Query("name")
	res, err := h.ItemService.GetItem(c.Context(), name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) UpdateItem(c *fiber.Ctx) error {
	var item UpdateReq
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	res, err := h.ItemService.UpdateItem(c.Context(), &item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) DeleteItem(c *fiber.Ctx) error {
	var item DelReq
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	err := h.ItemService.DeleteItem(c.Context(), &item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete item",
	})
}
