package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/saleh-ghazimoradi/GoInn/internal/repository"
	"github.com/saleh-ghazimoradi/GoInn/internal/service/service_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seeder interface {
	Seed(ctx context.Context) error
}

type seederService struct {
	userRepository  repository.UserRepository
	roomRepository  repository.RoomRepository
	hotelRepository repository.HotelRepository
}

func (s *seederService) Seed(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())

	users := GenerateUsers(10) // Generate 10 users
	for _, user := range users {
		_, err := s.userRepository.CreateUser(ctx, user)
		if err != nil {
			log.Println("Error while inserting user:", err)
			return err
		}
	}

	hotels := GenerateHotels(10) // Generate 10 hotels
	for _, hotel := range hotels {
		hotel, err := s.hotelRepository.InsertHotel(ctx, hotel)
		if err != nil {
			log.Println("Error while inserting hotel:", err)
			return err
		}

		rooms := GenerateRooms(rand.Intn(10)+5, hotel.Id) // 5 to 15 rooms per hotel
		for _, room := range rooms {
			room, err := s.roomRepository.InsertRoom(ctx, room)
			if err != nil {
				log.Println("Error while inserting room:", err)
				return err
			}
			hotel.Rooms = append(hotel.Rooms, room.Id)
		}

		_, err = s.hotelRepository.UpdateHotel(ctx, hotel)
		if err != nil {
			log.Println("Error updating hotel with rooms:", err)
			return err
		}
	}

	return nil
}

func NewSeedService(roomRepository repository.RoomRepository, hotelRepository repository.HotelRepository, userRepository repository.UserRepository) Seeder {
	return &seederService{
		userRepository:  userRepository,
		roomRepository:  roomRepository,
		hotelRepository: hotelRepository,
	}
}

func GenerateUsers(num int) []*service_models.User {
	firstNames := []string{"John", "Alice", "Michael", "Sophia", "David", "Emma", "James", "Olivia", "Robert", "Ava"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Miller", "Davis", "Garcia", "Rodriguez", "Wilson"}

	users := make([]*service_models.User, num)
	for i := 0; i < num; i++ {
		users[i] = &service_models.User{
			Id:        primitive.NewObjectID(),
			FirstName: firstNames[rand.Intn(len(firstNames))],
			LastName:  lastNames[rand.Intn(len(lastNames))],
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			Password:  "hashedpassword", // Replace with proper hashing
		}
	}
	return users
}

func GenerateRooms(num int, hotelID primitive.ObjectID) []*service_models.Room {
	roomTypes := []service_models.RoomType{
		service_models.SingleRoomType,
		service_models.DoubleRoomType,
		service_models.SeasideRoomType,
		service_models.DeluxeRoomType,
	}

	rooms := make([]*service_models.Room, num)
	for i := 0; i < num; i++ {
		roomType := roomTypes[rand.Intn(len(roomTypes))]
		basePrice := 50.0 + float64(rand.Intn(200))     // Base price between 50 - 250
		price := basePrice * (1.1 + rand.Float64()*0.5) // Final price with markup

		rooms[i] = &service_models.Room{
			Id:        primitive.NewObjectID(),
			Type:      roomType,
			BasePrice: basePrice,
			Price:     price,
			HotelId:   hotelID,
		}
	}
	return rooms
}

func GenerateHotels(num int) []*service_models.Hotel {
	hotelNames := []string{
		"Grand Palace Hotel", "Seaside Retreat", "Mountain View Lodge", "Urban Comfort Hotel",
		"Sunset Paradise", "Skyline Inn", "Lakeside Resort", "Royal Suites", "City Center Hotel", "Beachfront Haven",
	}
	locations := []string{
		"New York", "Los Angeles", "San Francisco", "Miami", "Chicago",
		"Seattle", "Las Vegas", "Austin", "Denver", "Orlando",
	}

	ratings := []int{
		1, 2, 3, 4, 5,
	}
	
	hotels := make([]*service_models.Hotel, num)
	for i := 0; i < num; i++ {
		hotels[i] = &service_models.Hotel{
			Id:       primitive.NewObjectID(),
			Name:     fmt.Sprintf("%s %d", hotelNames[i%len(hotelNames)], i+1),
			Location: locations[i%len(locations)],
			Rooms:    []primitive.ObjectID{},
			Rating:   ratings[i%len(ratings)],
		}
	}
	return hotels
}
