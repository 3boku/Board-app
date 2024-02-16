package service

import (
	"Back/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB, c *gin.Context) {
	id := c.Param("id") // JSON으로 전달된 id 값을 가져옵니다.

	var board entity.Board
	result := db.Where("id = ?", id).First(&board)

	if result.Error != nil {
		// 게시물을 찾을 수 없는 경우
		c.JSON(404, gin.H{"error": "게시물을 찾을 수 없습니다."})
		return
	}

	// 게시물을 삭제합니다.
	db.Delete(&board)

	c.JSON(200, gin.H{"message": "게시물이 삭제되었습니다."})
}
