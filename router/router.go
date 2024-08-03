package router

import (
	"github.com/DevloperAmanSingh/app-backend/controllers"
	"github.com/DevloperAmanSingh/app-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.Home(c)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.SignUp(c)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})

	app.Post("/events", func(c *fiber.Ctx) error {
		return handlers.AddEvent(c)
	})
	app.Get("/events", func(c *fiber.Ctx) error {
		return handlers.GetEvents(c)
	})
	app.Get("/events/nearby", func(c *fiber.Ctx) error {
		return handlers.GetNearbyEvents(c)
	})

	return app
}
