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

func GetEventByTag(c *fiber.Ctx) error {
	tag := c.Params("tag")
	if tag == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tag is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"tags": tag} // Assuming 'tags' is an array field in your event documents

	cursor, err := db.GetTripCollection().Find(ctx, filter)
	if err != nil {
		log.Printf("Error finding events by tag: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve events",
		})
	}
	defer cursor.Close(ctx)

	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		log.Printf("Error decoding events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to decode events",
		})
	}

	if len(events) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No events found for the given tag",
		})
	}

	return c.JSON(events)
}
