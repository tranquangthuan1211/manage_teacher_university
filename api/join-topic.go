package api

import (
	"myapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getJoinTopic(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", func(c *gin.Context) {
		var joinTopic []database.TopicBaseData
		err := db.Find(&joinTopic).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": joinTopic})
	})
}
