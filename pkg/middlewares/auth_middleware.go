package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dzakimaulana/golaundry/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// func Auth(ctx *fiber.Ctx) error {
// 	sess, err := session.Get
// 	token := sess.Get

// 	if token == "" {
// 		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "unauthorized",
// 		})
// 	}

// 	claims, err := utils.DecodeToken(token)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
// 			"error": "forbidden",
// 		})
// 	}
// 	ctx.Locals("userinfo", claims)
// 	return ctx.Next()
// }

func GetSession(session *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := session.Get(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		token := sess.Get("jwt")
		if token == "" || token == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}
		claims, err := utils.DecodeToken(token.(string))
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "forbidden",
			})
		}
		c.Locals("userinfo", claims)
		return c.Next()
	}
}

func OnlyAdmin(c *fiber.Ctx) error {
	userinfo := c.Locals("userinfo").(jwt.MapClaims)
	if userinfo["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden",
		})
	}
	return c.Next()
}
