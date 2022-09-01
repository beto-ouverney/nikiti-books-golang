package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Book presents a book
type Book struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Author   string             `json:"author" bson:"author"`
	Category []string           `json:"category" bson:"category"`
	Synopsis string             `json:"synopsis" bson:"synopsis"`
}
