package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/structs"
)

// MutantBody verifica que el body recibido sea válido
func MutantBody(c *fiber.Ctx) error {

	var sample structs.Sample
	var res structs.Response
	if err := sample.Unmarshal([]byte(c.Body())); err != nil {

		res.BadRequest("Body malformado.")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if err := sample.ValidateDNA(); err != nil {

		res.BadRequest(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	c.Context().SetUserValue("body", &sample)
	c.Next()
	return nil
}
