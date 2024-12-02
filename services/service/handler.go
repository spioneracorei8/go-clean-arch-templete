package auth

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}
