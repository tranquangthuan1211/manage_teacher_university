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
		err := db.Find(&users).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"numberTeacher": len(users),
			"data":          users,
		})
	})
}
func getTeachByID(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user database.Teach

		err := db.Where("MAGV = ?", id).Find(&user).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
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
func updateTeach(router *gin.RouterGroup, db *gorm.DB) {
	router.PUT("/:id", func(c *gin.Context) {
		var user database.Teach
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := db.Save(&user).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	})
}

func deleteTeach(router *gin.RouterGroup, db *gorm.DB) {
	router.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := db.Where("MAGV = ?", id).Delete(&database.Teach{}).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": id})
	})
}
