package api

import (
	"github.com/kalmecak/mutantfinder/api/mutant"
	"github.com/kalmecak/mutantfinder/api/stats"
	"github.com/kalmecak/mutantfinder/api/validate"

	"github.com/gofiber/fiber/v2"
)

// Router Es el que agrega las rutas a los paths del api v1
func Router(app fiber.Router) {

	app.Post("/mutant", validate.MutantBody, mutant.Post)
	app.Get("/stats", stats.Get)
}
