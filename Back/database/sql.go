package database

import (
	"Back/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func MySQLInit() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&entity.Board{})
	return db
}
