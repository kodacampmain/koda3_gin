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

// /products/order			GET	mengambil order aktif/keranjang
// /products/:productId		GET	mengambil product detail

// GET http://localhost:3000/products/order
// GET http://localhost:3000/products/1
