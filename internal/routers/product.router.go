package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/models"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool) {
	router.POST("/products", func(ctx *gin.Context) {
		var body models.Product
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		sql := "INSERT INTO products (name, promo_id, price) VALUES ($1,$2,$3) RETURNING id, name"
		values := []any{body.Name, body.PromoId, body.Price}
		var newProduct models.Product
		if err := db.QueryRow(ctx.Request.Context(), sql, values...).Scan(&newProduct.Id, &newProduct.Name); err != nil {
			log.Println("Internal Server Error: ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}
		// ctag, err := db.Exec(ctx.Request.Context(), sql, values...)
		// if err != nil {
		// 	log.Println("Internal Server Error: ", err.Error())
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 	})
		// 	return
		// }
		// if ctag.RowsAffected() == 0 {
		// 	ctx.JSON(http.StatusConflict, gin.H{
		// 		"success": false,
		// 	})
		// 	return
		// }
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    newProduct,
		})
	})
}
