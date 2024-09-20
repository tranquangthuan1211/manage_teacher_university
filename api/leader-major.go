package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getLeaderMajor(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", func(c *gin.Context) {
		var result []struct {
			Id          string `json:"MAGV" gorm:"column:MAGV"`
			Name        string `json:"HOTEN" gorm:"column:HOTEN"`
			Salary      string `json:"LUONG" gorm:"column:LUONG"`
			Sex         string `json:"PHAI" gorm:"column:PHAI"`
			DateOfBirth string `json:"NGSINH" gorm:"column:NGSINH"`
			Address     string `json:"DIACHI" gorm:"column:DIACHI"`
			MABM        string `json:"MABM"`
			TENBM       string `json:"TENBM"`
		}

		// Query the database for all users
		err := db.Table("BOMON").
			Select("GIAOVIEN.*, BOMON.MABM, BOMON.TENBM").
			Joins("JOIN GIAOVIEN ON GIAOVIEN.MAGV = BOMON.TRUONGBM").
			Find(&result).Error

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"numberLeader": len(result),
			"data":         result,
		})
	})

}

func updateLeaderMajor(router *gin.RouterGroup, db *gorm.DB) {
	router.PATCH("", func(c *gin.Context) {
		// Nhận dữ liệu từ body của request
		var input struct {
			MABM     string `json:"MABM"`
			TRUONGBM string `json:"TRUONGBM"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Dữ liệu không hợp lệ"})
			return
		}

		var existingMajor struct {
			MABM string
		}
		err := db.Table("BOMON").
			Select("MABM").
			Where("TRUONGBM = ?", input.TRUONGBM).
			First(&existingMajor).Error

		if err == nil {
			c.JSON(400, gin.H{
				"error": "Giáo viên này đã là trưởng bộ môn của bộ môn khác",
			})
			return
		}

		// Nếu giáo viên chưa là trưởng bộ môn, tiến hành cập nhật
		err = db.Table("BOMON").
			Where("MABM = ?", input.MABM).
			Update("TRUONGBM", input.TRUONGBM).Error

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Trả về thông báo thành công
		c.JSON(200, gin.H{
			"message": "Cập nhật trưởng bộ môn thành công",
		})
	})
}
