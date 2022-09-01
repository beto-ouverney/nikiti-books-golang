package entity

type Category struct {
	ID   int    `db:"category.id"`
	Name string `db:"category.name"`
}

type Author struct {
	ID   int    `db:"author.id"`
	Name string `db:"author.name"`
}

type Books struct {
	ID          uint64   `json:"id" db:"id"`
	Title       string   `json:"title" db:"title"`
	Author      Author   `json:"author" db:"author"`
	CategoryIDs []uint64 `json:"category"`
}

type BooksResponse struct {
	ID       uint64     `json:"id" db:"id"`
	Title    string     `json:"title" db:"title"`
	Author   Author     `json:"author" db:"author"`
	Category []Category `json:"category" db:"category"`
}
