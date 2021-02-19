package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Security maneja los middleware de seguridad y para optimizar las peticiones
func Security(app *fiber.App) {

	// Configuración de CORS
	app.Use(cors.New(cors.Config{
		// Permite todos los orígenes
		AllowOrigins: "*",
		// Sólo se deben agregar los vervos que se usan en el API
		AllowMethods: "POST,GET",
	}))

}
