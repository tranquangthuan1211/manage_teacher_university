package api

import (
	"myapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getTeach(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", func(c *gin.Context) {
		var users []database.Teach

		// Query the database for all users
		err := db.Where("MAGV like ?", "001").Find(&users).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": users})
	})
}
func createTeach(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("", func(c *gin.Context) {
		var user database.Teach
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Insert the user into the database
		err := db.Create(&user).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	})
}
