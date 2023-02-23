package main

import (
	"github.com/creztfallen/dnd-spells-api/configs"
	"github.com/creztfallen/dnd-spells-api/spells"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	spells.SpellRoute(app)

	app.Listen(":6000")
}
