package service

import (
	"Back/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Read(db *gorm.DB, c *gin.Context) {
	var boards []entity.Board
	db.Find(&boards)

	c.JSON(200, boards)
}
