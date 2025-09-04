package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
	"github.com/kodacampmain/koda3_gin/pkg"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	// menerima body
	var body models.StudentAuth
	if err := ctx.ShouldBind(&body); err != nil {
		if strings.Contains(err.Error(), "required") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Nama dan Password harus diisi",
			})
			return
		}
		if strings.Contains(err.Error(), "min") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Password minimum 4 karakter",
			})
			return
		}
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}

	// ambil data user
	student, err := a.ar.GetStudentWithPasswordAndRole(ctx.Request.Context(), body.Name)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Nama atau Password salah",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}

	// bandingkan password
	hc := pkg.NewHashConfig()
	isMatched, err := hc.CompareHashAndPassword(body.Password, student.Password)
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
		re := regexp.MustCompile("hash|crypto|argon2id|format")
		if re.Match([]byte(err.Error())) {
			log.Println("Error during Hashing")
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}

	if !isMatched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Nama atau Password salah",
		})
		return
	}
	// jika match, maka buatkan jwt dan kirim via response
	claims := pkg.NewJWTClaims(student.Id, student.Role)
	jwtToken, err := claims.GenToken()
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   jwtToken,
	})
}
