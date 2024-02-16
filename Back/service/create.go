package service

import (
	"Back/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Create(db *gorm.DB, c *gin.Context) {
	var board entity.Board

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.Create(&board).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "데이터베이스에 인서트를 완료했음",
	})
}
