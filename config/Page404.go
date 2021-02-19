package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/structs"
)

// Page404 maneja las peticiones cuyo path no son encontrados, redirige a una página
// random para que no sea analizada la respuesta
func Page404(app *fiber.App) {

	app.Use(func(c *fiber.Ctx) error {

		var redirect structs.RedirectList
		return c.Redirect(redirect.FunnyPage())
	})
}
