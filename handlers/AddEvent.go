package handlers

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/DevloperAmanSingh/app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateRandomID() (string, error) {
	// Generate a random number between 10000 and 99999
	nBig, err := rand.Int(rand.Reader, big.NewInt(90000))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%05d", nBig.Int64()+10000), nil
}

func AddEvent(c *fiber.Ctx) error {
	event := new(models.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	creator := c.Get("X-Creator-Name")
	if creator == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing creator name in header",
		})
	}

	event.ID = primitive.NewObjectID()
	event.Creator = creator

	// Generate a random 5-digit ID

	event.EventID, _ = generateRandomID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.GetTripCollection().InsertOne(ctx, event)
	if err != nil {
		log.Printf("Error inserting event: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}
