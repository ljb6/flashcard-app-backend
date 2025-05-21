package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
}