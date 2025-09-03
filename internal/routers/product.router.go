package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/handlers"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool) {
	productRouter := router.Group("/products")
	productRepository := repositories.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	productRouter.POST("/products", productHandler.AddNewProduct)
}
