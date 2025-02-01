package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateHotel struct {
	Name     string               `json:"name"`
	Location string               `json:"location"`
	Rooms    []primitive.ObjectID `json:"rooms"`
	Rating   int                  `json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeasideRoomType
	DeluxeRoomType
)

type CreateRoom struct {
	Type      RoomType           `json:"type"`
	BasePrice float64            `json:"basePrice"`
	Price     float64            `json:"price"`
	HotelId   primitive.ObjectID `json:"hotelId"`
}
