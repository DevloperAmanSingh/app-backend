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

// this is used to get all the events from the database

func GetEvents(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.GetTripCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error finding events: %v", err)
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

	return c.JSON(events)
}
