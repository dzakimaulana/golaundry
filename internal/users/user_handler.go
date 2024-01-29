package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Handler struct {
	UserService
	Session *session.Store
}

func NewHandler(s UserService, sess *session.Store) *Handler {
	return &Handler{
		UserService: s,
		Session:     sess,
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Check your username and password",
		})
	}
	sess, err := h.Session.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("jwt", res.AccessToken)
	sess.Save()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetAllUser(c *fiber.Ctx) error {
	res, err := h.UserService.GetAllUser(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.UserService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}

func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var req ResetPasswordReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}
	userinfo := c.Locals("userinfo").(jwt.MapClaims)
	id, ok := userinfo["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user ID in the JWT token",
		})
	}

	if err := h.UserService.ResetPassword(c.Context(), id, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success change password",
	})
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	sess, err := h.Session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error,
		})
	}
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout Success",
	})
}
