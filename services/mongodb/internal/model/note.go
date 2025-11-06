package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Note - модель заметки в MongoDB
type Note struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	Title     string        `bson:"title"`
	Body      string        `bson:"body"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt *time.Time    `bson:"updated_at,omitempty"`
}
