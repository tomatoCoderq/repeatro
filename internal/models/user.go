package models

import "time"

type User struct {
	UserId string `json:"userId" bson:"user_id"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	RegistrationDate time.Time `json:"registrationDate" bson:"registration_date"`
}