package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	ID      string `json:"id" gorm:"type:char(36);"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (board *Board) BeforeCreate(tx *gorm.DB) (err error) {
	board.ID = uuid.New().String()
	return
}
