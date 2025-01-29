package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateHotel struct {
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
}
