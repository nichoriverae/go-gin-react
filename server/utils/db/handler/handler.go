package handler

import (
	"fmt"

	// Models
	"go-gin-react/server/db/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// UpdatePost handles PUT to update a post
func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&post)
	db.Save(&post)
	c.JSON(200, post)
}

// CreatePost handles POST request to create new post
func CreatePost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	db.Create(&post)
	c.JSON(200, post)
}

// GetPostByID handles GET one post by ID
func GetPostByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var post models.Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

// GetAllPosts handle GET all posts from the db
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	if err := db.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, posts)
	}
}
