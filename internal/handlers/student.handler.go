package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
	"github.com/kodacampmain/koda3_gin/internal/utils"
	"github.com/kodacampmain/koda3_gin/pkg"
)

type StudentHandler struct {
	sr *repositories.StudentRepository
}

func NewStudentHandler(sr *repositories.StudentRepository) *StudentHandler {
	return &StudentHandler{
		sr: sr,
	}
}

func (s *StudentHandler) GetStudent(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}
	limit := 4
	offset := (page - 1) * limit

	students, err := s.sr.GetStudentData(ctx.Request.Context(), offset, limit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    students,
		})
		return
	}

	if len(students) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"data":    []any{},
			"page":    page,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    students,
		"page":    page,
	})
}

func (s *StudentHandler) EditImage(ctx *gin.Context) {
	// ambil dari form data
	var body models.StudentBody
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("Internal server error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal server error",
		})
		return
	}
	claims, _ := ctx.Get("claims")
	user, ok := claims.(pkg.Claims)
	if !ok {
		utils.HandleError(ctx, http.StatusInternalServerError, "Internal Server Error", "Cannot cast into pkg.claims")
		return
	}
	// proses menyimpan gambar di directory
	file := body.Image

	if file.Size > 2*1024 {
		// abort upload
	}

	ext := filepath.Ext(file.Filename)

	re := regexp.MustCompile("(png|jpg|jpeg|webp)$")
	if !re.Match([]byte(ext)) {
		// abort upload
	}

	filename := fmt.Sprintf("%d_images_%d%s", time.Now().UnixNano(), user.UserId, ext)
	location := filepath.Join("public", filename)
	if err := ctx.SaveUploadedFile(file, location); err != nil {
		log.Println("Upload Failed.\nCause: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Upload Failed",
		})
		return
	}
	// simpan ke database
	student, err := s.sr.EditImage(ctx.Request.Context(), filename, user.UserId)
	if err != nil {
		log.Println("Internal server error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    student,
	})
}

func (s *StudentHandler) GetStudentById(ctx *gin.Context) {
	// ambil id dari claims/token
	claims, _ := ctx.Get("claims")
	user, _ := claims.(pkg.Claims)

	student, err := s.sr.GetStudentById(ctx.Request.Context(), user.UserId)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	utils.HandleResponse(ctx, http.StatusOK, models.StudentResponse{
		SuccessResponse: models.SuccessResponse{
			Success: true,
			Status:  200,
		},
		Data: student,
	})
}
