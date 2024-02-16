package controller

import (
	"Back/database"
	"Back/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(port string) {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	r.Use(cors.New(config))

	db := database.MySQLInit()

	r.POST("/create", func(c *gin.Context) {
		service.Create(db, c)
	})

	r.GET("/posts", func(c *gin.Context) {
		service.Read(db, c)
	})

	r.DELETE("/delete/:id", func(c *gin.Context) {
		service.Delete(db, c)
	})

	r.PUT("/update/:id", func(c *gin.Context) {
		service.Update(db, c)
	})
	r.Run(port)
}
