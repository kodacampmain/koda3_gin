package main

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// inisialisasi engine gin
	router := gin.Default()
	router.Use(MyLogger)
	router.Use(CORSMiddleware)

	// config := cors.Config{
	// 	AllowOrigins: []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"},
	// 	AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders: []string{"Authorization", "Content-Type"},
	// }

	// router.Use(cors.New(config))

	// setup routing
	router.GET("/ping", func(ctx *gin.Context) {
		requestId := ctx.GetHeader("X-Request-ID")
		contentType := ctx.GetHeader("Content-Type")
		ctx.JSON(http.StatusOK, gin.H{
			"message":     "pong",
			"requestId":   requestId,
			"contentType": contentType,
		})
	})
	router.GET("/ping/:id/:param2", func(ctx *gin.Context) {
		pingId := ctx.Param("id")
		param2 := ctx.Param("param2")
		q := ctx.Query("q")
		log.Printf("%s, %s, %s\n", pingId, param2, q)
		ctx.JSON(http.StatusOK, gin.H{
			"param":  pingId,
			"param2": param2,
			"q":      q,
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
		log.Println(body)
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

func MyLogger(ctx *gin.Context) {
	log.Println("Start")
	start := time.Now()
	ctx.Next() // Next digunakan untuk lanjut ke middleware/handler berikutnya
	duration := time.Since(start)
	log.Printf("Durasi Request: %dus\n", duration.Microseconds())
}

func CORSMiddleware(ctx *gin.Context) {
	// if ctx.Request.Method == http.MethodPost {
	// 	ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	// 	return
	// }
	// memasangkan header-header CORS
	// setup whitelist origin
	whitelist := []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"}
	origin := ctx.GetHeader("Origin")
	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}
	// header untuk preflight cors
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	// tangani apabila bertemu preflight
	if ctx.Request.Method == http.MethodOptions {
		// ctx.Header("X-DEBUG", "preflight-handled")
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	// ctx.Header("X-DEBUG", "actual request")
	ctx.Next()
}

type Response struct {
	Message string
	Status  string
}

type Body struct {
	Id      int    `json:"id" binding:"required"`
	Message string `json:"msg"`
	Gender  string `json:"gender"`
}

func ValidateBody(body Body) error {
	if body.Id <= 0 {
		return errors.New("id harus diatas 0")
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
