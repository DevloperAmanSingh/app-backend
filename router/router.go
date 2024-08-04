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
	app.Get("/events/tag/all", func(c *fiber.Ctx) error {
		return handlers.GetEvents(c)
	})
	app.Get("/events/tag/:tag", func(c *fiber.Ctx) error {
		return handlers.GetEventByTag(c)
	})
	app.Get("/events/nearby", func(c *fiber.Ctx) error {
		return handlers.GetNearbyEvents(c)
	})
	app.Post("/addBookmark", func(c *fiber.Ctx) error {
		return handlers.AddBookmark(c)
	})
	app.Get("/getBookmarks/:username", func(c *fiber.Ctx) error {
		return handlers.GetBookmarks(c)
	})
	app.Delete("/removeBookmark", func(c *fiber.Ctx) error {
		return handlers.RemoveBookmark(c)
	})

	return app
}
