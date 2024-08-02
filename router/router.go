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

	app.Post("/createTrip", func(c *fiber.Ctx) error {
		return handlers.CreateTrip(c)
	})
	app.Get("/getTripInfo/:id", func(c *fiber.Ctx) error {
		return handlers.GetTripInfo(c)
	})
	app.Get("/getTrips", func(c *fiber.Ctx) error {
		return handlers.GetAllTrips(c)
	})

	return app
}
