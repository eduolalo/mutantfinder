package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/storage"
)

// Get maneja las paticiones get del paquete stats
func Get(c *fiber.Ctx) error {

	stats := storage.GetStats()

	return c.Status(fiber.StatusOK).JSON(stats)
}
