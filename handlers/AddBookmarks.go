package handlers

import (
	"context"
	"log"
	"time"

	db "github.com/DevloperAmanSingh/app-backend/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddBookmark(c *fiber.Ctx) error {
	type BookmarkRequest struct {
		Username string `json:"username"`
		EventID  string `json:"event_id"`
	}

	req := new(BookmarkRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": req.Username}
	update := bson.M{"$addToSet": bson.M{"bookmarks": req.EventID}}

	_, err := db.GetUserCollection().UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Printf("Error adding bookmark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add bookmark",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bookmark added successfully",
	})
}
