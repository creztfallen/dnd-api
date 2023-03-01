package dndclasses

import "github.com/gofiber/fiber/v2"

func ClassRoute(app *fiber.App) {
	app.Post("/api/classes", CreateClass)
	app.Get("/api/classes", getAllClasses)
}
