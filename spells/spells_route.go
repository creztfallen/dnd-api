package spells

import "github.com/gofiber/fiber/v2"

func SpellRoute(app *fiber.App) {
	app.Post("/spells", CreateSpell)
}
