package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kalmecak/mutantfinder/api"
	"github.com/kalmecak/mutantfinder/config"
	"github.com/kalmecak/mutantfinder/environment"
)

func main() {

	/**************************************************************************/
	/*               Verifcación de las variables de entorno                  */
	/**************************************************************************/
	if ok := environment.Validate(); ok == false {

		log.Println("Entorno no configurado correctamente")
		return
	}
	log.Println("Environment is ok!")

	app := fiber.New(fiber.Config{
		CaseSensitive:    true,
		ServerHeader:     "Mutant Finder",
		DisableKeepalive: true,
	})

	app.Static("/", "./public")
	// Configuraciones
	config.Accepts(app)
	config.Security(app)
	config.General(app)

	// router
	api.Router(app)
	// Manejador de Páginas no encontradas
	config.Page404(app)

	port := os.Getenv("PORT")
	if port == "" {

		port = "3000"
	}
	app.Listen(":" + port)
}
