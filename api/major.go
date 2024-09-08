package api

import (
	"myapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getMajor(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", func(c *gin.Context) {
		var majors []database.Major
		err := db.Find(&majors).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": majors})
	})
}
