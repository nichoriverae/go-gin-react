package models

// Post model for the database
type Post struct {
	ID      int64  `json: "id"`
	Title   string `json: "title"`
	Content string `json: "title"`
}
