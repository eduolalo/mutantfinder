package mutant

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/storage"
	"github.com/kalmecak/mutantfinder/structs"
)

// Post majena los post al paquete mutant
func Post(c *fiber.Ctx) error {

	sample := c.Context().UserValue("body").(*structs.Sample)
	var res structs.Response

	sample.Analyze()

	var status int
	mutant := sample.IsMutant()

	if mutant {

		res.Ok("")
		status = fiber.StatusOK
	} else {

		res.Forbridden("")
		status = fiber.StatusForbidden
	}
	// Almacenar los resultados
	go storage.StoreResult(mutant, *sample)

	return c.Status(status).JSON(res)
}
