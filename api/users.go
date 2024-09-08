package api

import (
	"myapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loginHandler(router *gin.RouterGroup, loginHandler gin.HandlerFunc) {
	router.POST("/users/login", loginHandler)
}

func registerHandler(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/users/register", func(c *gin.Context) {
		req := &database.Register{}

		// Bind the JSON payload to the request struct
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}

		// Attempt to create a new user in the database
		result := db.Table("USERS").Create(req)
		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Failed to create user",
			})
			return
		}

		// Return a success response
		c.JSON(200, gin.H{
			"message": "User registered successfully",
		})
	})
}
