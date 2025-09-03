package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
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
