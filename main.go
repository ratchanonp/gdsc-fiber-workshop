package main

import (
	"workshop/project/controller"
	"workshop/project/database"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/books", controller.GetBooks)
	app.Get("/books/:id", controller.GetBook)
	app.Post("/books", controller.NewBook)
	app.Delete("/books/:id", controller.DeleteBook)
	app.Put("/books/:id", controller.UpdateBook)
}

func main() {
	app := fiber.New()

	database.InitDB()
	setUpRoutes(app)

	app.Listen(":3000")
}
