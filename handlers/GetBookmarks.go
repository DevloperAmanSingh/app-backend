package handlers

import (
	"context"
	"log"
	"time"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/DevloperAmanSingh/app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBookmarks(c *fiber.Ctx) error {
	username := c.Params("username")

	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username is required",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch user bookmarks
	userFilter := bson.M{"username": username}
	var user models.User
	err := db.GetUserCollection().FindOne(ctx, userFilter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		log.Printf("Error finding user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
		})
	}

	// Fetch event details
	if len(user.Bookmarks) == 0 {
		return c.JSON([]models.Event{})
	}

	eventFilter := bson.M{"event_id": bson.M{"$in": user.Bookmarks}}
	cursor, err := db.GetTripCollection().Find(ctx, eventFilter)
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
