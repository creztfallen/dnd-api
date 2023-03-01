package main

import (
	"github.com/dnd-api/configs"
	"github.com/dnd-api/dnd-classes"
	"github.com/dnd-api/dnd-spells"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	spells.SpellRoute(app)
	dndclasses.ClassRoute(app)
	app.Listen(":6000")
}
