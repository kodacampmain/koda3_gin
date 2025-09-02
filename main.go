package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi engine gin
	router := gin.Default()
	// setup routing
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Message: "Rute Salah",
			Status:  "Rute Tidak Ditemukan",
		})
	})
	// jalankan engine gin
	router.Run("localhost:3000")
}

type Response struct {
	Message string
	Status  string
}
