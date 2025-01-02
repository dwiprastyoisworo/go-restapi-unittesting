package dto

import (
	"go-restapi-unittesting/app/helpers"
	"time"
)

type Users struct {
	//ID          string    `bson:"_id"`
	Email       string    `bson:"email"`
	Name        string    `bson:"name"`
	Password    string    `bson:"password"`
	PhoneNumber string    `bson:"phone_number"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	DeletedAt   time.Time `bson:"deleted_at"`
}

func (u Users) setTime() {
	currentDate := time.Now().In(helpers.GenerateTimeLocation())
	u.CreatedAt = currentDate
	u.UpdatedAt = currentDate
}

type CreateUserPayload struct {
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number"`
}

func (p CreateUserPayload) UserPayloadToUsers() Users {
	users := Users{
		Name:        p.Name,
		Email:       p.Email,
		PhoneNumber: p.PhoneNumber,
	}
	users.setTime()
	return users
}
