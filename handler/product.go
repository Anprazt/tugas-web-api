package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ProductHandler(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	price := c.Query("price")
	time := time.Now()

	c.JSON(http.StatusOK, gin.H{
		"Message":      "Item Name",
		"Id":           id,
		"Product Name": name,
		"Price":        price,
		"Time":         time,
	})
}
