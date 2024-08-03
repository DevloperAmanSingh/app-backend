package handlers

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/DevloperAmanSingh/app-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	rand.Seed(time.Now().UnixNano())
	event.EventID = strconv.Itoa(rand.Intn(90000) + 10000)

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
