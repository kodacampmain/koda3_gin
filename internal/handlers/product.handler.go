package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
)

type ProductHandler struct {
	pr *repositories.ProductRepository
}

func NewProductHandler(pr *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		pr: pr,
	}
}

func (p *ProductHandler) AddNewProduct(ctx *gin.Context) {
	var body models.Product
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	newProduct, err := p.pr.AddNewProduct(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	// ctag, err := p.pr.InsertNewProduct(ctx.Request.Context(), body)
	// if err != nil {
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
}
