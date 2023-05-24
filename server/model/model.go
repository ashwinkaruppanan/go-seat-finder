package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllotRequest struct {
	Departments []string `json:"departments"`
	Rooms       []int    `json:"rooms"`
	Date        int64    `json:"date"`
}

type FindSeatRequest struct {
	RollNumber string `json:"roll_number"`
	// Date       int64  `json:"date"`
}

type AllotSeat struct {
	ID         primitive.ObjectID `bson:"_id"`
	RollNumber string             `bson:"roll_number" json:"roll_number"`
	RoomNumber int                `bson:"room_number" json:"room_number"`
	SeatNumber int                `bson:"seat_number" json:"seat_number"`
	Date       int64              `bson:"date" json:"date"`
}
