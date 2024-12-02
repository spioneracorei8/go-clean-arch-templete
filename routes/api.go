package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	app *fiber.App
}

func NewRoute(app *fiber.App) Route {
	return Route{app: app}
}
