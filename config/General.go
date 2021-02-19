package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// General maneja los middleware de seguridad y para optimizar las peticiones
func General(app *fiber.App) {

	// Usar la máxima compresión
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Log de los requests
	format := "--> ${method} ${path} <--\n ${time} - HTTPCode: ${status} - IP: ${ips} ${ip} - Latencia: ${latency}\n"
	app.Use(logger.New(logger.Config{
		Format:     format,
		TimeFormat: "02-Jan-2006",
	}))

	// Set la respuesta en Content-Type application/json
	app.Use(func(c *fiber.Ctx) error {

		c.Accepts("application/json")
		return c.Next()
	})

}
