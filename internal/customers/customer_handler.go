package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct {
	CustomerService
}

func NewHandler(s CustomerService) *Handler {
	return &Handler{
		CustomerService: s,
	}
}

func (h *Handler) AddNewCustomer(c *fiber.Ctx) error {
	var cs CustomerReq
	if err := c.BodyParser(&cs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	res, err := h.CustomerService.AddNewCustomer(c.Context(), &cs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetCustomerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	csId, err := uuid.Parse(id)
	if id == "" || err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id parameter not detected",
		})
	}
	cs := &CusIdReq{
		ID: csId,
	}
	res, err := h.CustomerService.GetCustomerByID(c.Context(), cs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
