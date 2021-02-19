package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/structs"
)

// Accepts agraga configuraciones para las de los requests
func Accepts(app *fiber.App) {

	// Sólamente se aceptarán los Content-Type application/json para recibir datos
	app.Use(func(c *fiber.Ctx) error {

		if c.Method() == "GET" {
			return c.Next()
		}

		ct := c.Context().Request.Header.ContentType()
		if !strings.Contains(string(ct), "application/json") {

			var res structs.Response
			res.BadRequest("Content-Type erroneo.")
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		return c.Next()
	})

}
