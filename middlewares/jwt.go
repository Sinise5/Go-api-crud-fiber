package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

func JWTMiddleware() fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte("secret"), // Ganti dengan key yang lebih aman
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}
