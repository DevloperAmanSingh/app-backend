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

func GetTripInfo(c *fiber.Ctx) error {
	tripID := c.Params("id")
	if tripID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing trip ID",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var trip models.Trip
	err := db.GetTripCollection().FindOne(ctx, bson.M{"tripID": tripID}).Decode(&trip)
	if err != nil {
		log.Printf("Error getting trip: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get trip",
		})
	}

	return c.Status(fiber.StatusOK).JSON(trip)
}
