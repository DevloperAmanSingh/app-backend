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

func GetNearbyEvents(c *fiber.Ctx) error {
	type LocationQuery struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Radius    float64 `json:"radius"` // in meters
	}

	locationQuery := new(LocationQuery)
	if err := c.BodyParser(locationQuery); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{locationQuery.Longitude, locationQuery.Latitude},
				},
				"$maxDistance": locationQuery.Radius,
			},
		},
	}

	cursor, err := db.GetTripCollection().Find(ctx, filter)
	if err != nil {
		log.Printf("Error finding nearby events: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve nearby events",
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
