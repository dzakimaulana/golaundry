package middlewares

import (
	"github.com/dzakimaulana/golaundry/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden",
		})
	}
	ctx.Locals("userinfo", claims)
	return ctx.Next()
}
