package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/handlers"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
	"github.com/redis/go-redis/v9"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	productRouter := router.Group("/products")
	productRepository := repositories.NewProductRepository(db, rdb)
	productHandler := handlers.NewProductHandler(productRepository)

	productRouter.POST("/", productHandler.AddNewProduct)
	productRouter.PATCH("/:productId", productHandler.EditProduct)
	productRouter.GET("/", productHandler.GetProducts)
}
