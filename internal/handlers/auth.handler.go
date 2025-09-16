package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
	"github.com/kodacampmain/koda3_gin/internal/utils"
	"github.com/kodacampmain/koda3_gin/pkg"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

// Login
// @tags	auth
// @router	/auth [POST]
// @accept json
// @param body body models.StudentAuth true "email and password input"
// @produce json
// @success 200 {object} models.AuthResponse
// @failure 400 {object} models.BadRequestError
// @failure 500 {object} models.InternalServerError
func (a *AuthHandler) Login(ctx *gin.Context) {
	// menerima body
	var body models.StudentAuth
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "required") {
			utils.HandleError(ctx, http.StatusBadRequest, "Nama dan Password harus diisi", "Bad Request")
			return
		}
		if strings.Contains(err.Error(), "min") {
			utils.HandleError(ctx, http.StatusBadRequest, "Password minimum 4 karakter", "Bad Request")
			return
		}
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	// ambil data user
	student, err := a.ar.GetStudentWithPasswordAndRole(ctx.Request.Context(), body.Name)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			utils.HandleError(ctx, http.StatusBadRequest, "Nama atau Password salah", "Bad Request")
			return
		}
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	// bandingkan password
	hc := pkg.NewHashConfig()
	isMatched, err := hc.CompareHashAndPassword(body.Password, student.Password)
	if err != nil {
		re := regexp.MustCompile("hash|crypto|argon2id|format")
		if re.Match([]byte(err.Error())) {
			log.Println("Error during Hashing")
		}
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	if !isMatched {
		utils.HandleError(ctx, http.StatusBadRequest, "Nama atau Password salah", "Bad Request")
		return
	}
	// jika match, maka buatkan jwt dan kirim via response
	claims := pkg.NewJWTClaims(student.Id, student.Role)
	jwtToken, err := claims.GenToken()
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}
	utils.HandleResponse(ctx, http.StatusOK, models.AuthResponse{
		SuccessResponse: models.SuccessResponse{
			Success: true,
			Status:  200,
		},
		Data: models.AuthData{Token: jwtToken},
	})
}
