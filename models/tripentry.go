package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trip struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TripID      string             `bson:"tripID"`
	TripName    string             `bson:"tripName"`
	TeamMembers []string           `bson:"teamMembers"`
	TotalBudget int                `bson:"totalBudget"`
	FromDate    string             `bson:"fromDate"`
	ToDate      string             `bson:"toDate"`
	Place       string             `bson:"place"`
	Creator     string             `bson:"creator"`
}
