package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
)

type Reservation struct {
	Id                primitive.ObjectID `bson:"_id"`
	AccommodationID   string             `bson:"accommodation_id"`
	GuestNum          uint               `bson:"guest_num"`
	Start             date.Date          `bson:"start"`
	End               date.Date          `bson:"end"`
	ReservationStatus ReservationStatus  `bson:"reservation_status"`
	UserID            string             `bson:"user_id"`
}
