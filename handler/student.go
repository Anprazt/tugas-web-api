package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MahasiswaInput struct {
	Nama  string `json:"Nama" binding:"required"`
	Prodi string `json:"Prodi" binding:"required"`
}

func MahasiswaHandler(c *gin.Context) {
	var mahasiswaInput MahasiswaInput
	time := time.Now()

	err := c.ShouldBindJSON(&mahasiswaInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Field %s harus diisi. Pesan: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message":        "Berhasil input data",
		"Nama Mahasiswa": mahasiswaInput.Nama,
		"Prodi":          mahasiswaInput.Prodi,
		"Time":           time,
	})
}
