package models

import (
	"github.com/jinzhu/gorm"
)

// BlogPost model for the database
type BlogPost struct {
	gorm.Model
	// ID      int64  `json: "id"`
	Title   string `json: "title"`
	Content string `json: "content"`
}
