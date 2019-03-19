package main

import (
	"log"
	"os"
	"net/http"
	"github.com/joho/godotenv"

	// DB packages - external
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Gin packages
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	// DB - internal

	"go-gin-react/server/db/models"
	"go-gin-react/server/utils/db/handler"
	// "go-gin-react/server/utils/db/migrate"
)

var db *gorm.DB
var err error

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnStr := os.Getenv("MYSQLSTR")
	
	// Set up DB connection
	db, err = gorm.Open("mysql", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.BlogPost{})
	//migrate.Start(db)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../client", true)))

	// API Route Groups
	api := router.Group("/api/")
	{

		// Blog Subgroup
		blog := api.Group("/blog/")
		{
			blog.GET("/", handler.GetAllBlogPosts(db))

			blog.GET("/:id", handler.GetBlogPostByID(db))

			blog.PUT("/:id", handler.UpdateBlogPost(db))

			blog.POST("/", handler.CreateBlogPost(db))

			blog.DELETE("/:id", handler.DeleteBlogPost(db))
		}
		// Begin api base routes
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "base route reached",
			})
		})
		// more api routes go here
		// example:
		// api.GET("/specific/path/to/route/:id", someHandler)
	}
	// Start and run the server
	router.Run(":5000")
}

// // A standard handler function
// func someHandler(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")
// 	c.JSON(http.StatusOK, JSON_content)
// }
