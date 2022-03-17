package main

import (
	"github.com/Anprazt/tugas-web-api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/home", handler.HomeHandler)
	r.GET("/item", handler.QueryStringItem)
	r.GET("/id/:id", handler.ProductHandler)
	r.Run(":8080")
}
