package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kodacampmain/koda3_gin/internal/models"
	"github.com/kodacampmain/koda3_gin/internal/repositories"
	"github.com/kodacampmain/koda3_gin/internal/utils"
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

func (p *ProductHandler) EditProduct(ctx *gin.Context) {
	var body models.EditProductBody
	if err := ctx.ShouldBind(&body); err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}

	productId, _ := strconv.Atoi(ctx.Param("productId"))
	product, err := p.pr.EditProduct(ctx.Request.Context(), body, productId)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}
	utils.HandleResponse(ctx, http.StatusOK, models.ProductResponse{
		SuccessResponse: models.SuccessResponse{
			Success: true,
			Status:  200,
		},
		Data: product,
	})
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	data, err := p.pr.GetProducts(ctx.Request.Context())
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error(), "Internal Server Error")
		return
	}
	utils.HandleResponse(ctx, http.StatusOK, models.ProductsResponse{
		SuccessResponse: models.SuccessResponse{
			Success: true,
			Status:  200,
		},
		Data: data,
	})
}
