package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/middlewares"
	"github.com/kodacampmain/koda3_gin/internal/utils"
	"github.com/redis/go-redis/v9"

	docs "github.com/kodacampmain/koda3_gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.MyLogger)
	router.Use(middlewares.CORSMiddleware)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/wekekwek/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// config := cors.Config{
	// 	AllowOrigins: []string{"http://127.0.0.1:5500", "http://127.0.0.1:3001"},
	// 	AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders: []string{"Authorization", "Content-Type"},
	// }

	// router.Use(cors.New(config))

	router.Static("/img", "public")

	// setup routing
	InitPingRouter(router)
	InitStudentRouter(router, db)
	InitProductRouter(router, db, rdb)
	InitAuthRouter(router, db)

	// catch all route
	router.NoRoute(func(ctx *gin.Context) {
		utils.HandleError(ctx, http.StatusNotFound, "rute tidak ditemukan", "404 Route")
	})

	return router
}
