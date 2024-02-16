package service

import (
	"Back/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(db *gorm.DB, c *gin.Context) {
	id := c.Param("id") // JSON으로 전달된 id 값을 가져옵니다.

	var board entity.Board
	result := db.Where("id = ?", id).First(&board)

	if result.Error != nil {
		// 게시물을 찾을 수 없는 경우
		c.JSON(404, gin.H{"error": "게시물을 찾을 수 없습니다."})
		return
	}

	// JSON으로 전달된 요청 본문(body) 값을 파싱하여 게시물 필드를 업데이트합니다.
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(400, gin.H{"error": "잘못된 요청 형식입니다."})
		return
	}

	// 게시물을 저장하여 업데이트합니다.
	db.Save(&board)

	c.JSON(200, gin.H{"message": "게시물이 수정되었습니다."})
}
