package handler

import (
	"fmt"

	// Models
	"go-gin-react/server/db/models"

	// Gin and Gorm
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UpdateBlogPost handles PUT to update a blogPost
func UpdateBlogPost(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var blogPost models.BlogPost
		id := c.Params.ByName("id")
		if err := db.Where("id = ?", id).First(&blogPost).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.BindJSON(&blogPost)
		db.Save(&blogPost)
		c.JSON(200, blogPost)
	}
}

// DeleteBlogPost handles DELETE request
func DeleteBlogPost(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var blogPost models.BlogPost
		id := c.Params.ByName("id")
		fmt.Printf("The id requested is %s \n", id)
		if err := db.Where("id = ?", id).Delete(&blogPost).Error; err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			msg := fmt.Sprintf("blog post %s has been deleted", id)
			c.JSON(200, gin.H{"message": msg})
		}
	}
}
	
// CreateBlogPost handles POST request to create new blogPost
func CreateBlogPost(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var blogPost models.BlogPost
		c.BindJSON(&blogPost)
		db.Create(&blogPost)
		c.JSON(200, blogPost)
	}
}

// GetBlogPostByID handles GET one blogPost by ID
func GetBlogPostByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		var blogPost models.BlogPost
		if err := db.Where("id = ?", id).First(&blogPost).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
				c.JSON(200, blogPost)
		}
	}
}
		
// GetAllBlogPosts handle GET all posts from the db
func GetAllBlogPosts(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var blogPosts []models.BlogPost
		if err := db.Find(&blogPosts).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, blogPosts)
		}
	}
}
		