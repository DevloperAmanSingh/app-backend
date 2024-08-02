package handlers

import (
	"context"
	"log"
	"time"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/DevloperAmanSingh/app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTrips(c *fiber.Ctx) error {
	creator := c.Get("X-Creator-Name")
	if creator == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing creator name in header",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var trips []models.Trip
	cursor, err := db.GetTripCollection().Find(ctx, bson.M{"creator": creator})
	if err != nil {
		log.Printf("Error finding trips: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve trips",
		})
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &trips); err != nil {
		log.Printf("Error decoding trips: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user trips",
		})
	}

	return c.Status(fiber.StatusOK).JSON(trips)
}
