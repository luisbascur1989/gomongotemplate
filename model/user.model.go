package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	CreatedAt time.Time          `json:"-" bson:"created_at"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
}
