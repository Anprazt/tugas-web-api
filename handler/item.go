package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ItemHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "item barang",
		"id":      id,
	})
}

func QueryStringItem(c *gin.Context) {
	judul := c.Query("judul")
	harga := c.Query("harga")
	waktuSekarang := time.Now()

	c.JSON(http.StatusOK, gin.H{
		"Message":   "Item Buku",
		"Item Name": judul,
		"Waktu":     waktuSekarang,
		"Price":     harga,
	})
}
