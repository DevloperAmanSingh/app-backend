package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	EventID     string             `json:"event_id" bson:"event_id"`
	EventName   string             `json:"event_name" bson:"event_name"`
	Description string             `json:"description" bson:"description"`
	Location    GeoJSONPoint       `json:"location" bson:"location"`
	StartTime   time.Time          `json:"start_time" bson:"start_time"`
	EndTime     time.Time          `json:"end_time" bson:"end_time"`
	Creator     string             `json:"creator" bson:"creator"`
}

type GeoJSONPoint struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
