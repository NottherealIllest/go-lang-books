package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Author string `bson:"author,omitempty" json:"author,omitempty"`
	Country string `bson:"country,omitempty" json:"country,omitempty"`
	Image string `bson:"image,omitempty" json:"image_url,omitempty"`
	Language string `bson:"language,omitempty" json:"language,omitempty"`
	Link string `bson:"link,omitempty" json:"link,omitempty"`
	Pages int `bson:"pages,omitempty" json:"pages,omitempty"`
	Title string `bson:"title,omitempty" json:"title,omitempty"`
	Year string `bson:"year,omitempty" json:"year,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
