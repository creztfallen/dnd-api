package main

import (
	"github.com/creztfallen/dnd-spells-api/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Listen(":6000")
}
