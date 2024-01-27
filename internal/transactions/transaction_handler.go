package transactions

import "github.com/gofiber/fiber/v2"

type Handler struct {
	TransactionService
}

func NewHandler(s TransactionService) *Handler {
	return &Handler{
		TransactionService: s,
	}
}

func (h *Handler) AddTs(c *fiber.Ctx) error {
	var ts TransactionReq
	if err := c.BodyParser(&ts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	res, err := h.TransactionService.AddTs(c.Context(), &ts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetAllTs(c *fiber.Ctx) error {
	res, err := h.TransactionService.GetAllTs(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetTsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.TransactionService.GetTsById(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
