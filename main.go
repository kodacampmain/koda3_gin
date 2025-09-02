package main

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi engine gin
	router := gin.Default()
	// router.Use(func(ctx *gin.Context) {
	// 	log.Println("Middleware")
	// 	ctx.Next()
	// })

	// setup routing
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/ping", func(ctx *gin.Context) {
		body := Body{}
		// data-binding, memasukkan body ke dalam variabel golang
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		if err := ValidateBody(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// log.Println(body)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"body":    body,
		})
	})

	// catch all route
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

type Body struct {
	Id      int    `json:"id"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}

func ValidateBody(body Body) error {
	if body.Id <= 0 {
		return errors.New("id tidak boleh dibawah 0")
	}
	if len(body.Message) < 8 {
		return errors.New("panjang pesan harus diatas 8 karakter")
	}
	re, err := regexp.Compile("^[lLpPmMfF]$")
	if err != nil {
		return err
	}
	if !re.Match([]byte(body.Gender)) {
		return errors.New("gender harus berisikan huruf l, L, m, M, f, F, p, P")
	}
	return nil
}
