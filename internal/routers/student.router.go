package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/handlers"
	"github.com/kodacampmain/koda3_gin/internal/middlewares"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
)

func InitStudentRouter(router *gin.Engine, db *pgxpool.Pool) {
	studentRouter := router.Group("/students")
	sr := repositories.NewStudentRepository(db)
	sh := handlers.NewStudentHandler(sr)

	studentRouter.GET("", middlewares.VerifyToken, sh.GetStudent)
}
