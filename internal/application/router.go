package application

import "github.com/gofiber/fiber/v2"

func (app *App) loadRoutes() {
	// hello check
	app.fiberApp.Get("/", func(c *fiber.Ctx) error {
		var result int
		err := app.db.Get(&result, "SELECT 1")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"hello":    "world",
			"database": "available",
		})
	})
}
