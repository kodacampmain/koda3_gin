package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/handlers"
)

func InitPingRouter(router *gin.Engine) {
	pingRouter := router.Group("/ping")
	ph := handlers.NewPingHandler()

	pingRouter.GET("", ph.GetPing)
	pingRouter.GET("/:id/:param2", ph.GetPingWithParam)
	pingRouter.POST("", ph.PostPing)
}
