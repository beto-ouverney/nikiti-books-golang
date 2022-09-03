package entity

// Book presents a book
type Book struct {
	Title    string   `json:"title" bson:"title"`
	Author   string   `json:"author" bson:"author"`
	Category []string `json:"category" bson:"category"`
	Synopsis string   `json:"synopsis" bson:"synopsis"`
}
