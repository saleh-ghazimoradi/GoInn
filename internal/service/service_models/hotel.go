package service_models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	Id       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeasideRoomType
	DeluxeRoomType
)

type Room struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type      RoomType           `json:"type" bson:"type"`
	BasePrice float64            `json:"basePrice" bson:"basePrice"`
	Price     float64            `json:"price" bson:"price"`
	HotelId   primitive.ObjectID `json:"hotelId" bson:"hotelId"`
}
